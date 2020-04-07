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
	"time"

	"github.com/djlauk/punchcard/data"
	"github.com/spf13/cobra"
)

type summaryEntry struct {
	Project *data.Project
	Hours   float64
}

// reportSummaryCmd represents the reportSummary command
var reportSummaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Print sum of hours per project and grand total",
	Run: func(cmd *cobra.Command, args []string) {
		strStart, err := cmd.Flags().GetString("start")
		if err != nil {
			log.Fatal(err)
		}
		strEnd, err := cmd.Flags().GetString("end")
		if err != nil {
			log.Fatal(err)
		}

		pcd := readData()
		if strStart != "" && strEnd != "" {
			start := parseDate(strStart)
			end := parseDate(strEnd)
			summaryEntries := make(map[string]summaryEntry, 0)
			for name := range pcd.Projects {
				p := pcd.Projects[name]
				summaryEntries[name] = summaryEntry{
					Project: &p,
					Hours:   0,
				}
			}
			for _, entry := range pcd.Entries {
				if entry.Start.Before(start) || entry.Start.Equal(end) || entry.Start.After(end) {
					continue
				}
				s, ok := summaryEntries[entry.Project]
				if !ok {
					log.Fatalf("Inconsistency! Project from log not in projects: %s", entry.Project)
				}
				s.Hours += entry.End.Local().Sub(entry.Start.Local()).Hours()
				summaryEntries[entry.Project] = s
			}
			printSummary(&start, &end, summaryEntries)
		}
	},
}

func printSummary(start *time.Time, end *time.Time, entries map[string]summaryEntry) {
	fmt.Printf("SUMMARY FOR %s - %s\n\n", formatDate(start), formatDate(end))
	fmt.Println(`"Project";"Reference";"Hours";"% total"`)
	grandTotal := 0.0
	for _, entry := range entries {
		grandTotal += entry.Hours
	}
	keys := make([]string, 0, len(entries))
	for k := range entries {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		entry := entries[k]
		fmt.Printf("\"%s\";\"%s\";\"%.2f\";\"%.1f\"\n", entry.Project.Name, entry.Project.Reference, entry.Hours, entry.Hours*100.0/grandTotal)
	}
	fmt.Printf("\nTOTAL: %.2fh\n", grandTotal)
}

func init() {
	reportCmd.AddCommand(reportSummaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportSummaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportSummaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	reportSummaryCmd.Flags().String("start", "today", "Start date of summary (inclusive)")
	reportSummaryCmd.Flags().String("end", "tomorrow", "End date of summary (exclusive)")
}
