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
	"time"

	"github.com/djlauk/punchcard/data"
	"github.com/spf13/cobra"
)

// workStatusCmd represents the status command
var workStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display what you are working on right now",
	Long:  `Displays what project you are working on, and what note you left.`,
	Run: func(cmd *cobra.Command, args []string) {
		pcd := readData()
		PrintStatus(pcd.Current, true)
	},
}

// PrintStatus will print the currently (incomplete)
func PrintStatus(e *data.WorkLogEntry, verbose bool) {
	//loc := time.Local
	if e == nil {
		if verbose {
			fmt.Println("Not working on anything right now")
		}
		return
	}
	fmt.Printf(`Started on:         %s (%s ago)
Project:            %s
Message:            %s
Reference:          %s
`, e.Start.Local().Format(DateTimeFormat), time.Since(e.Start.Local()).Round(time.Second), e.Project, e.Message, e.Reference)
}

func init() {
	workCmd.AddCommand(workStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
