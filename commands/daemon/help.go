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

package daemon

import (
	"context"
	"fmt"
	"log"
	"os"
)

func Help(l *log.Logger, c context.Context, a []string) {
	message := `Usage:
	
	threshold daemon:<subcommand> [arguments]
	
Available subcommands and arguments:

	start [device] [threshold (MiB)] [interval (Min)] # Start service
	stop [device]                                     # Stop service
	status [device]                                   # View service status
	help                                              # View help message
	
Example:

	threshold daemon:start eth0 1024 1`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
