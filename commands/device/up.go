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
	"strings"

	"github.com/vishvananda/netlink"
)

func Up(l *log.Logger, c context.Context, a []string) {
	if len(a) != 1 {
		Help(l, c, a)
		return
	}

	device, err := netlink.LinkByName(a[0])

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	if device.Attrs().OperState == netlink.OperUp {
		l.Printf("[INFO] the device is already up\n")
		fmt.Fprintf(os.Stdout, "the device is already up\n")

		return
	}

	err = netlink.LinkSetUp(device)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	l.Printf("[INFO] the device is up\n")
	fmt.Fprintf(os.Stdout, "the device is up\n")
}
