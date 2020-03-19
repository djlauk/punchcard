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

var project string
var message string
var reference string

func entryFromFlags(cmd *cobra.Command) *data.WorkLogEntry {
	entry := data.WorkLogEntry{}
	var err error
	if entry.Reference, err = cmd.Flags().GetString("reference"); err != nil {
		log.Fatal(err)
	}
	if entry.Message, err = cmd.Flags().GetString("message"); err != nil {
		log.Fatal(err)
	}
	if entry.Project, err = cmd.Flags().GetString("project"); err != nil {
		log.Fatal(err)
	}
	str, err := cmd.Flags().GetString("time")
	if err != nil {
		log.Fatal(err)
	}
	entry.Start = parseTime(str)
	return &entry
}

// workStartCmd represents the workStart command
var workStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start recording work on a project",
	Run: func(cmd *cobra.Command, args []string) {
		entry := entryFromFlags(cmd)
		pcd := readData()
		if pcd.Current != nil {
			if err := pcd.FinishCurrent(entry.Start); err != nil {
				log.Fatal(err)
			}
		}
		if err := pcd.SetCurrent(entry); err != nil {
			log.Fatal(err)
		}
		writeData(pcd)
	},
}

func init() {
	workCmd.AddCommand(workStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workStartCmd.Flags().StringP("project", "p", "", "Name of the project")
	workStartCmd.Flags().StringP("message", "m", "", "Log message")
	workStartCmd.Flags().StringP("reference", "r", "", "External reference (e.g. URL, bug id, ...)")
	workStartCmd.Flags().StringP("time", "t", "now", "Time when work started")

	workStartCmd.MarkFlagRequired("project")
	workStartCmd.MarkFlagRequired("message")

}
