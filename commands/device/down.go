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
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vishvananda/netlink"
)

func Down(l *log.Logger, c context.Context, a []string) {
	if len(a) != 2 {
		Help(l, c, a)
		return
	}

	device, err := netlink.LinkByName(a[0])

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	if device.Attrs().OperState == netlink.OperDown {
		l.Printf("[INFO] the device is already down\n")
		fmt.Fprintf(os.Stderr, "the device is already down\n")

		return
	}

	threshold, err := strconv.ParseFloat(a[1], 64)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	file, err := os.Open("/proc/net/dev")

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	defer file.Close()

	var bandwidth float64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || !strings.HasPrefix(line, device.Attrs().Name) {
			continue
		}

		values := strings.Fields(line)

		download, err := strconv.ParseFloat(values[1], 64)

		if err != nil {
			l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

			return
		}

		upload, err := strconv.ParseFloat(values[9], 64)

		if err != nil {
			l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

			return
		}

		bandwidth = (download + upload) / (1024 * 1024)
	}

	if bandwidth <= threshold {
		l.Printf("[INFO] the bandwidth doesn't meet the threshold yet\n")
		fmt.Fprintf(os.Stdout, "the bandwidth doesn't meet the threshold yet\n")

		return
	}

	err = netlink.LinkSetDown(device)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	l.Printf("[INFO] the device is down\n")
	fmt.Fprintf(os.Stdout, "the device is down\n")
}
