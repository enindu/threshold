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
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/vishvananda/netlink"
)

func Status(l *log.Logger, c context.Context, a []string) {
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

	serviceProperties, err := connection.GetUnitPropertiesContext(c, serviceFileName)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	serviceFileState := fmt.Sprintf("exists in %s", serviceFilePath)

	_, err = os.Stat(serviceFilePath)

	if err != nil {
		serviceFileState = "not exists"
	}

	timerProperties, err := connection.GetUnitPropertiesContext(c, timerFileName)

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	timerFileState := fmt.Sprintf("exists in %s", timerFilePath)

	_, err = os.Stat(timerFilePath)

	if err != nil {
		timerFileState = "not exists"
	}

	message := serviceFileName + `:

	Active state:    ` + serviceProperties["ActiveState"].(string) + `
	Unit file state: ` + serviceProperties["UnitFileState"].(string) + `
	File state:      ` + serviceFileState + `
	
` + timerFileName + `:

	Active state:    ` + timerProperties["ActiveState"].(string) + `
	Unit file state: ` + timerProperties["UnitFileState"].(string) + `
	File state:      ` + timerFileState + `
	
Use "systemctl status ` + serviceFileName + `" and "systemctl status ` + timerFileName + `" for more information.`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
