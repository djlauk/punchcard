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

// projectRenameCmd represents the projectRename command
var projectRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a project",
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}
		newName, err := cmd.Flags().GetString("new-name")
		if err != nil {
			log.Fatal(err)
		}
		pcd := readData()
		if err := pcd.RenameProject(project, newName); err != nil {
			log.Fatal(err)
		}
		writeData(pcd)
	},
}

func init() {
	projectCmd.AddCommand(projectRenameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectRenameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectRenameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectRenameCmd.Flags().String("project", "", "Name of the project to rename")
	projectRenameCmd.Flags().String("new-name", "", "New name of the project")

	projectRenameCmd.MarkFlagRequired("project")
	projectRenameCmd.MarkFlagRequired("new-name")
}
