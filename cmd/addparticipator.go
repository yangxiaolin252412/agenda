/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

var name, title string

// addparticipatorCmd represents the addparticipator command
var addparticipatorCmd = &cobra.Command{
	Use:   "addparticipator -n [participator name] -t [meeting title]",
	Short: "add a participator to a meeting with title",
	Long:  `Input command model like : addparticipator -n Golang -t BeautifulGo`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_t, _ := cmd.Flags().GetString("title")
		tmp_n, _ := cmd.Flags().GetString("name")
		if service.GetFlag() == true {
			service.Addparticipator(tmp_n, tmp_t)
		} else {
			fmt.Println("You don't log in!")
		}

	},
}

func init() {
	rootCmd.AddCommand(addparticipatorCmd)
	addparticipatorCmd.Flags().StringP("name", "n", "", "participator name")
	addparticipatorCmd.Flags().StringP("title", "t", "", "meeting title")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
