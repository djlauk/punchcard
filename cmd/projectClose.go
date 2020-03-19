/*
Copyright © 2020 Daniel J. Lauk <daniel.lauk@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// projectCloseCmd represents the projectClose command
var projectCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close a project to stop accidental recording of time",
	Long: `You can close a project in order to prevent accidental
recording of further work logs.`,
	Run: func(cmd *cobra.Command, args []string) {
		pcd := readData()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		p, ok := pcd.Projects[name]
		if !ok {
			log.Fatalf("Project not found: %s", name)
		}
		if pcd.Current != nil && pcd.Current.Project == name {
			log.Fatalf("Cannot close project; it's currently being worked on")
		}
		p.Closed = true
		pcd.Projects[name] = p
		writeData(pcd)
	},
}

func init() {
	projectCmd.AddCommand(projectCloseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCloseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCloseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectCloseCmd.Flags().StringP("name", "n", "", "Name of the project")

	projectCloseCmd.MarkFlagRequired("name")
}
