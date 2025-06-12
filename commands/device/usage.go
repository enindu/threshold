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

func Usage(l *log.Logger, c context.Context, a []string) {
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

	file, err := os.Open("/proc/net/dev")

	if err != nil {
		l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

		return
	}

	defer file.Close()

	var (
		download float64
		upload   float64
		total    float64
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || !strings.HasPrefix(line, device.Attrs().Name) {
			continue
		}

		values := strings.Fields(line)

		download, err = strconv.ParseFloat(values[1], 64)

		if err != nil {
			l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

			return
		}

		upload, err = strconv.ParseFloat(values[9], 64)

		if err != nil {
			l.Printf("[ERROR] %s\n", strings.ToLower(err.Error()))
			fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))

			return
		}

		total = (download + upload) / (1024 * 1024)
	}

	message := device.Attrs().Name + `:
	
	Download: ` + fmt.Sprintf("%.3f", download/(1024*1024)) + ` MiB
	Upload:   ` + fmt.Sprintf("%.3f", upload/(1024*1024)) + ` MiB
	Total:    ` + fmt.Sprintf("%.3f", total) + ` MiB`

	fmt.Fprintf(os.Stdout, "%s\n", message)
}
