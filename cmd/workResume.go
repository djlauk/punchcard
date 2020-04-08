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

// workResumeCmd represents the workResume command
var workResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick up where you stopped",
	Run: func(cmd *cobra.Command, args []string) {
		str, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal(err)
		}
		t := parseTime(str)
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			log.Fatal(err)
		}
		finishCurrent, err := cmd.Flags().GetBool("finish-current")
		if err != nil {
			log.Fatal(err)
		}
		pcd := readData()
		var entry *data.WorkLogEntry
		if project != "" {
			entry = pcd.GetLastEntryForProject(project)
		} else {
			entry = pcd.GetLastEntry()
		}
		if entry == nil {
			log.Fatal("Could not find entry to resume")
		}
		newEntry := *entry
		newEntry.Start = t
		newEntry.End = zeroTime()
		if pcd.Current != nil && finishCurrent {
			pcd.FinishCurrent(newEntry.Start)
		}
		if err := pcd.SetCurrent(&newEntry); err != nil {
			log.Fatal(err)
		}
		writeData(pcd)
	},
}

func init() {
	workCmd.AddCommand(workResumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workResumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workResumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workResumeCmd.Flags().StringP("time", "t", "now", "Time when work was resumed")
	workResumeCmd.Flags().StringP("project", "p", "", "Project to resume")
	workResumeCmd.Flags().BoolP("finish-current", "f", false, "Finish current work and then resume")
}
