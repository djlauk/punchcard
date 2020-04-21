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

	"github.com/djlauk/punchcard/data"
	"github.com/spf13/cobra"
)

// reportListCmd represents the reportList command
var reportListCmd = &cobra.Command{
	Use:   "list",
	Short: "List entries from the work log",
	Run: func(cmd *cobra.Command, args []string) {
		strStart, err := cmd.Flags().GetString("start")
		if err != nil {
			log.Fatal(err)
		}
		strEnd, err := cmd.Flags().GetString("end")
		if err != nil {
			log.Fatal(err)
		}
		showAll, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}
		last, err := cmd.Flags().GetInt("last")
		if err != nil {
			log.Fatal(err)
		}

		pcd := readData()
		printEntryHeader()
		if showAll {
			last = len(pcd.Entries)
		}
		if last > 0 {
			if last > len(pcd.Entries) {
				last = len(pcd.Entries)
			}
			for i := last; i > 0; i-- {
				idx := len(pcd.Entries) - i
				entry := pcd.Entries[idx]
				printEntry(idx, &entry)
			}
		} else {
			start := parseDate(strStart)
			end := parseDate(strEnd)
			for idx, entry := range pcd.Entries {
				if entry.Start.Before(start) || entry.Start.Equal(end) || entry.Start.After(end) {
					continue
				}
				printEntry(idx, &entry)
			}
		}
	},
}

func printEntryHeader() {
	fmt.Println(`"Id";"Start";"End";"Hours";"Project";"Message";"Reference"`)
}

func printEntry(id int, entry *data.WorkLogEntry) {
	fmt.Printf("\"%d\";\"%s\";\"%s\";\"%.2f\";\"%s\";\"%s\";\"%s\"\n", id, formatDateTime(&entry.Start), formatDateTime(&entry.End), entry.End.Local().Sub(entry.Start.Local()).Hours(), entry.Project, entry.Message, entry.Reference)
}

func init() {
	reportCmd.AddCommand(reportListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	reportListCmd.Flags().String("start", "today", "Start of list (inclusive)")
	reportListCmd.Flags().String("end", "tomorrow", "End of list (exclusive)")
	reportListCmd.Flags().Int("last", 0, "Show last NUM entries")
	reportListCmd.Flags().BoolP("all", "a", false, "Show ALL entries")
}
