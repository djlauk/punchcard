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
	"fmt"
	"log"
	"sort"

	"github.com/spf13/cobra"
)

// projectListCmd represents the projectList command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		withClosed, err := cmd.Flags().GetBool("with-closed")
		if err != nil {
			log.Fatal(err)
		}
		pcd := readData()
		keys := make([]string, 0, len(pcd.Projects))
		for k := range pcd.Projects {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		fmt.Printf("\"Name\";\"Description\";\"Reference\";\"Closed\"\n")
		for _, k := range keys {
			p := pcd.Projects[k]
			if p.Closed {
				if withClosed {
					fmt.Printf("\"%s\";\"%s\";\"%s\";\"%v\"\n", p.Name, p.Description, p.Reference, p.Closed)
				}
				continue
			}
			fmt.Printf("\"%s\";\"%s\";\"%s\";\"%v\"\n", p.Name, p.Description, p.Reference, p.Closed)
		}
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectListCmd.Flags().Bool("with-closed", false, "Also list closed projects")
}
