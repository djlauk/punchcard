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

// workStopCmd represents the workStop command
var workStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Finish working",
	Long:  `Finish what you're currently working at. Commits the temporary entry to the log of work entries.`,
	Run: func(cmd *cobra.Command, args []string) {
		str, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal(err)
		}
		t := parseTime(str)
		pcd := readData()
		if err := pcd.FinishCurrent(t); err != nil {
			log.Fatal(err)
		}
		writeData(pcd)
	},
}

func init() {
	workCmd.AddCommand(workStopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workStopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workStopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workStopCmd.Flags().StringP("time", "t", "now", "Time when work stopped")
}
