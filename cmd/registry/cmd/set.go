// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/apigee/registry/cmd/registry/tools"
	"github.com/apigee/registry/connection"
	"github.com/apigee/registry/gapic"
	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/names"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

var setFilter string
var setLabelID string

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVar(&setFilter, "filter", "", "Filter selected resources")
	setCmd.Flags().StringVar(&setLabelID, "label_id", "", "Label to set on selected resources")
}

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set labels and properties on matching entities.",
	Long:  "set labels and properties on matching entities.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		client, err := connection.NewClient(ctx)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		// Initialize task queue.
		taskQueue := make(chan tools.Task, 1024)
		workerCount := 64
		for i := 0; i < workerCount; i++ {
			tools.WaitGroup().Add(1)
			go tools.Worker(ctx, taskQueue)
		}

		err = matchAndHandleSetCmd(ctx, client, taskQueue, args[0])
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				log.Fatalf("%s", err.Error())
			} else {
				log.Fatalf("%s", st.Message())
			}
		}

		close(taskQueue)
		tools.WaitGroup().Wait()
	},
}

type setTask struct {
	ctx          context.Context
	client       connection.Client
	resourceName string
	resourceKind string
}

func (task *setTask) Run() error {
	if setLabelID != "" {
		log.Printf("setting %s/labels/%s", task.resourceName, setLabelID)
		return tools.SetLabel(task.ctx, task.client, &rpc.Label{
			Subject: task.resourceName,
			Label:   setLabelID,
		})
	} else {
		return nil
	}
}

func matchAndHandleSetCmd(
	ctx context.Context,
	client connection.Client,
	taskQueue chan tools.Task,
	name string,
) error {
	if m := names.ProjectRegexp().FindStringSubmatch(name); m != nil {
		return setProjects(ctx, client, m, setFilter, taskQueue)
	} else if m := names.ApiRegexp().FindStringSubmatch(name); m != nil {
		return setAPIs(ctx, client, m, setFilter, taskQueue)
	} else if m := names.VersionRegexp().FindStringSubmatch(name); m != nil {
		return setVersions(ctx, client, m, setFilter, taskQueue)
	} else if m := names.SpecRegexp().FindStringSubmatch(name); m != nil {
		return setSpecs(ctx, client, m, setFilter, taskQueue)
	} else {
		return fmt.Errorf("unsupported resource name.")
	}
}

func setProjects(
	ctx context.Context,
	client *gapic.RegistryClient,
	segments []string,
	filterFlag string,
	taskQueue chan tools.Task) error {
	return tools.ListProjects(ctx, client, segments, filterFlag, func(project *rpc.Project) {
		taskQueue <- &setTask{
			ctx:          ctx,
			client:       client,
			resourceName: project.Name,
			resourceKind: "project",
		}
	})
}

func setAPIs(
	ctx context.Context,
	client *gapic.RegistryClient,
	segments []string,
	filterFlag string,
	taskQueue chan tools.Task) error {
	return tools.ListAPIs(ctx, client, segments, filterFlag, func(api *rpc.Api) {
		taskQueue <- &setTask{
			ctx:          ctx,
			client:       client,
			resourceName: api.Name,
			resourceKind: "api",
		}
	})
}

func setVersions(
	ctx context.Context,
	client *gapic.RegistryClient,
	segments []string,
	filterFlag string,
	taskQueue chan tools.Task) error {
	return tools.ListVersions(ctx, client, segments, filterFlag, func(version *rpc.Version) {
		taskQueue <- &setTask{
			ctx:          ctx,
			client:       client,
			resourceName: version.Name,
			resourceKind: "version",
		}
	})
}

func setSpecs(
	ctx context.Context,
	client *gapic.RegistryClient,
	segments []string,
	filterFlag string,
	taskQueue chan tools.Task) error {
	return tools.ListSpecs(ctx, client, segments, filterFlag, func(spec *rpc.Spec) {
		taskQueue <- &setTask{
			ctx:          ctx,
			client:       client,
			resourceName: spec.Name,
			resourceKind: "spec",
		}
	})
}