// Copyright © 2019 NAME HERE adam.placzek@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/brudnyhenry/harrayp/cmd"
	"github.com/brudnyhenry/harrayp/config"
	"github.com/spf13/cobra"
)

var (
	// shaman provides the shaman cli/server functionality
	harraypTool = &cobra.Command{
		Use:               "harrayp",
		Short:             "harrayp - cli tool for hp p2000 array",
		Long:              ``,
		PersistentPreRunE: readConfig,
	}
)

func init() {
	harraypTool.AddCommand(cmd.GetCmd)

}

func readConfig(ccmd *cobra.Command, args []string) error {
	if err := config.LoadConfigFile(); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return err
	}
	return nil
}

func main() {
	harraypTool.Execute()
}
