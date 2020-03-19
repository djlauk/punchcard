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

	"github.com/spf13/cobra"
)

var examplesText = `These are common example commands for use with punchcard:

# add a project so you can track time on it
punchcard project add --name 'skynet' --description 'ultimately powerful AI, what could go wrong?!'
punchcard project add --name 'email' --description 'non-project-related email'

# working on stuff
punchcard work start --project skynet --note 'Design app architecture'
# start again to change to a different thing
punchcard work start --project email --note 'Sift through inbox'
punchcard work status
punchcard work stop
# pick up where you left (using the same project and note)
# just like: punchcard work start --project email --note 'Sift through inbox'
punchcard work resume
# go back to work on the last task of project skynet
# just like: punchcard work start --project skynet --note 'Design app architecture'
punchcard work resume --project skynet

# report a list of recorded times
punchcard report list --start 2020-03-01 --end 2020-03-02
punchcard report list --start 2020-03-01 --end 2020-04-01
punchcard report list --last 10

# report a summary: hours per project and total
punchcard report summary --start 2020-03-01 --end 2020-04-01`

// examplesCmd represents the examples command
var examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Show common examples of how to use punchcard",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(examplesText)
	},
}

func init() {
	rootCmd.AddCommand(examplesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// examplesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// examplesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
