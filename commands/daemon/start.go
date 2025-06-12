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
	"os/exec"
	"strconv"
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/vishvananda/netlink"
)

func Start(l *log.Logger, c context.Context, a []string) {
	if len(a) != 3 {
		Help(l, c, a)
		return
	}

	device, err := netlink.LinkByName(a[0])

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	threshold, err := strconv.ParseFloat(a[1], 64)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	if threshold < 1 {
		l.Printf("[ERROR] %v\n", errInvalidThreshold)
		fmt.Fprintf(os.Stderr, "%v\n", errInvalidThreshold)

		return
	}

	interval, err := strconv.ParseInt(a[2], 10, 0)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	if interval < 1 {
		l.Printf("[ERROR] %v\n", errInvalidInterval)
		fmt.Fprintf(os.Stderr, "%v\n", errInvalidInterval)

		return
	}

	serviceFileName := fmt.Sprintf("%s-threshold.service", device.Attrs().Name)
	serviceFilePath := fmt.Sprintf("/etc/systemd/system/%s", serviceFileName)
	timerFileName := fmt.Sprintf("%s-threshold.timer", device.Attrs().Name)
	timerFilePath := fmt.Sprintf("/etc/systemd/system/%s", timerFileName)

	connection, err := dbus.NewSystemdConnectionContext(c)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	defer connection.Close()

	executable, err := exec.LookPath("threshold")

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	serviceContent := `[Unit]
Description=Automatically monitor network usage and disable the network interface when a specified threshold is exceeded

[Service]
Type=oneshot
ExecStart=` + executable + ` device:down ` + device.Attrs().Name + ` ` + fmt.Sprintf("%.3f", threshold)

	serviceFile, err := os.OpenFile(serviceFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	defer serviceFile.Close()

	_, err = serviceFile.WriteString(serviceContent)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	timerContent := `[Unit]
Description=Run ` + serviceFileName + ` periodically

[Timer]
OnBootSec=1min
OnUnitActiveSec=` + fmt.Sprintf("%d", interval) + `min
Unit=` + serviceFileName + `

[Install]
WantedBy=timers.target`

	timerFile, err := os.OpenFile(timerFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	defer timerFile.Close()

	_, err = timerFile.WriteString(timerContent)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	_, err = connection.StartUnitContext(c, timerFileName, "replace", nil)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	l.Printf("[INFO] the daemon is started\n")
	fmt.Fprintf(os.Stdout, "the daemon is started\n")
}
