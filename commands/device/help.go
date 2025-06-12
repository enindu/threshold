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

package device

import (
	"context"
	"fmt"
	"log"
	"os"
)

func Help(l *log.Logger, c context.Context, a []string) {
	message := `Usage:
	
	threshold device:<subcommand> [arguments]
	
Available subcommands and arguments:

	up [device]                     # Enable device
	down [device] [threshold (mib)] # Disable device
	usage [device]                  # View usage
	help                            # View help message
	
Example:

	threshold device:up eth0`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
