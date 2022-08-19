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

package config

import (
	"fmt"

	"github.com/apigee/registry/cmd/registry/cmd/config/configurations"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Set, view, and unset properties used by Registry CLI.",
	}

	cmd.AddCommand(configurations.Command())
	cmd.AddCommand(getCommand())
	cmd.AddCommand(listCommand())
	cmd.AddCommand(setCommand())
	cmd.AddCommand(unsetCommand())
	return cmd
}

type UnknownPropertyError struct {
	property string
}

func (n UnknownPropertyError) Error() string {
	return fmt.Sprintf("Unknown property: %q.", n.property)
}