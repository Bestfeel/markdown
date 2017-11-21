// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/Bestfeel/markdown/server"
)

var (
	globalAddr = ":8080"
	globalPath = "."
	globalCss  = "github"
)
var RootCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Powerful markdown online",
	Long:  `This is a powerful online tool about markdown that allows you to run a service online to view markdown files,and can also serve as a static server`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		server.RunMarkDownServer(globalAddr, globalPath, globalCss)

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().StringVarP(&globalAddr, "addr", "a", ":8080", "server address")
	RootCmd.Flags().StringVarP(&globalPath, "path", "p", ".", "sever path")
	RootCmd.Flags().StringVarP(&globalCss, "css", "c", "github", "markdown for css style.example [github|mou|marxico]")

}
