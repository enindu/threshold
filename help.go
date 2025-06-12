// This file is part of Threshold.
//
// Threshold is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// Threshold is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// Threshold. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
)

func help() {
	message := `Usage:

	threshold <command>:<subcommand> [arguments]

Available commands:

	daemon
	device

Use "threshold <command>:help" to see more information.`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
