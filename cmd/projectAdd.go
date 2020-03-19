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

	"github.com/djlauk/punchcard/data"
	"github.com/spf13/cobra"
)

func projectFromFlags(cmd *cobra.Command) *data.Project {
	ref, err := cmd.Flags().GetString("reference")
	if err != nil {
		log.Fatal(err)
	}
	desc, err := cmd.Flags().GetString("description")
	if err != nil {
		log.Fatal(err)
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatal(err)
	}
	p := data.Project{
		Name:        name,
		Description: desc,
		Reference:   ref,
		Closed:      false,
	}
	return &p
}

// projectAddCmd represents the projectAdd command
var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add another project",
	Run: func(cmd *cobra.Command, args []string) {
		p := projectFromFlags(cmd)
		pcd := readData()
		if err := pcd.AddProject(p); err != nil {
			log.Fatal(err)
		}
		writeData(pcd)
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectAddCmd.Flags().String("name", "", "Name of the project")
	projectAddCmd.Flags().String("description", "", "Description of the project")
	projectAddCmd.Flags().String("reference", "", "External reference of the project")

	projectAddCmd.MarkFlagRequired("name")
}
