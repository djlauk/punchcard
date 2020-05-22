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

	"github.com/spf13/cobra"
)

type dailyTotalEntry struct {
	Date  string
	Hours float64
}

// reportDailyTotals represents the reportDailyTotals command
var reportDailyTotalsCmd = &cobra.Command{
	Use:   "daily-totals",
	Short: "Print daily totals for a sequence of days and a grand total of all days together",
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
			reportEntries := make(map[string]dailyTotalEntry, 0)

			for _, entry := range pcd.Entries {
				if entry.Start.Before(start) || entry.Start.Equal(end) || entry.Start.After(end) {
					continue
				}
				dateStart := formatDate(&entry.Start)
				dateEnd := formatDate(&entry.End)
				if dateStart != dateEnd {
					// split into 2 entries
					t := entry.End.Local()
					midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
					hoursDay2 := entry.End.Local().Sub(midnight).Hours()
					hoursDay1 := entry.End.Local().Sub(entry.Start.Local()).Hours() - hoursDay2
					addToDailyTotal(reportEntries, dateStart, hoursDay1)
					addToDailyTotal(reportEntries, dateEnd, hoursDay2)
				} else {
					addToDailyTotal(reportEntries, dateStart, entry.End.Local().Sub(entry.Start.Local()).Hours())
				}
			}
			printDailyTotals(&start, &end, reportEntries)
		}
	},
}

func addToDailyTotal(entries map[string]dailyTotalEntry, date string, hours float64) {
	entry, ok := entries[date]
	if !ok {
		entry = dailyTotalEntry{
			Date:  date,
			Hours: 0,
		}
	}
	entry.Hours += hours
	entries[date] = entry
}

func printDailyTotals(start *time.Time, end *time.Time, entries map[string]dailyTotalEntry) {
	fmt.Printf("DAILY TOTALS FOR %s - %s\n\n", formatDate(start), formatDate(end))
	fmt.Println(`"Date";"Hours"`)
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
		fmt.Printf("\"%s\";\"%.2f\"\n", entry.Date, entry.Hours)
	}
	fmt.Printf("\nTOTAL: %.2fh\n", grandTotal)
}

func init() {
	reportCmd.AddCommand(reportDailyTotalsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportSummaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportSummaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	reportDailyTotalsCmd.Flags().String("start", "today", "Start date of report (inclusive)")
	reportDailyTotalsCmd.Flags().String("end", "tomorrow", "End date of report (exclusive)")
}
