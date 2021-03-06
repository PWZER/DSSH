/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/PWZER/dssh/config"
	//"github.com/PWZER/dssh/ssh"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <host> <local_src_path> <remote_dest_path>",
	Short: "upload local files to remote host",
	Long:  "upload local files to remote host",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		switch len(args) {
		case 0:
			return config.GetHostNames(), cobra.ShellCompDirectiveNoFileComp
		case 1:
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		case 2:
			return []string{}, cobra.ShellCompDirectiveNoFileComp
		default:
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
