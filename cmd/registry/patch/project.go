// Copyright 2022 Google LLC. All Rights Reserved.
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

package patch

import (
	"context"
	"fmt"
	"os"

	"github.com/apigee/registry/cmd/registry/core"
	"github.com/apigee/registry/gapic"
	"github.com/apigee/registry/log"
	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/registry/names"
)

// ExportProject writes a project as a YAML file.
func ExportProject(ctx context.Context, client *gapic.RegistryClient, adminClient *gapic.AdminClient, name string) {
	projectName, err := names.ParseProject(name)
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("Failed to parse project name")
	}
	apisDir := "apis"
	err = os.MkdirAll(apisDir, 0777)
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("Failed to create output directory")
	}
	core.ListAPIs(ctx, client, projectName.Api(""), "", func(message *rpc.Api) {
		bytes, header, err := exportAPI(ctx, client, message)
		if err != nil {
			log.FromContext(ctx).WithError(err).Fatal("Failed to export artifact")
		}
		filename := fmt.Sprintf("%s/%s.yaml", apisDir, header.Metadata.Name)
		err = os.WriteFile(filename, bytes, 0644)
		if err != nil {
			log.FromContext(ctx).WithError(err).Fatal("Failed to write output YAML")
		}
	})
	artifactsDir := "artifacts"
	err = os.MkdirAll(artifactsDir, 0777)
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("Failed to create output directory")
	}
	core.ListArtifacts(ctx, client, projectName.Artifact(""), "", false, func(message *rpc.Artifact) {
		bytes, header, err := exportArtifact(ctx, client, message)
		if err != nil {
			log.FromContext(ctx).WithError(err).Fatal("Failed to export artifact")
		}
		filename := fmt.Sprintf("%s/%s.yaml", artifactsDir, header.Metadata.Name)
		err = os.WriteFile(filename, bytes, 0644)
		if err != nil {
			log.FromContext(ctx).WithError(err).Fatal("Failed to write output YAML")
		}
	})
}
