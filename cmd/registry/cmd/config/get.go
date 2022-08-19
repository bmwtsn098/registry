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

	"github.com/apigee/registry/pkg/config"
	"github.com/spf13/cobra"
)

func getCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get PROPERTY",
		Short: "Print the value of a property.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			_, c, err := config.ActiveRaw()
			if err != nil {
				if err == config.NoActiveConfigurationError {
					return fmt.Errorf(`No active configuration. Use 'registry config configurations' to manage.`)
				}
				return fmt.Errorf("Cannot read config: %v", err)
			}

			if !contains(c.Properties(), args[0]) {
				return UnknownPropertyError{args[0]}
			}

			m, err := c.FlatMap()
			if err != nil {
				return fmt.Errorf("Cannot decode config: %v", err)
			}

			cmd.Println(m[args[0]])
			return nil
		},
	}
	return cmd
}