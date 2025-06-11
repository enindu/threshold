// Threshold is a simple application based on systemd that automatically brings
// down a network interface after a specified amount of data has passed through
// it.
// Copyright (C) 2025  Enindu Alahapperuma
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/enindu/threshold/commands/daemon"
	"github.com/enindu/threshold/commands/device"
)

func main() {
	account, err := user.Current()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	if account.Uid != "0" {
		fmt.Fprintf(os.Stderr, "%v\n", errNonRoot)
		return
	}

	file, err := os.OpenFile("/var/log/threshold.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToLower(err.Error()))
		return
	}

	defer file.Close()

	logger := log.New(file, log.Prefix(), log.Flags())

	dispatchers := map[string]func(*log.Logger, context.Context, []string){
		"daemon:start": daemon.Start,
		"daemon:stop":  daemon.Stop,
		"daemon:help":  daemon.Help,
		"device:up":    device.Up,
		"device:down":  device.Down,
		"device:help":  device.Help,
	}

	inputs := os.Args

	if len(inputs) < 2 {
		help()
		return
	}

	instruction := inputs[1]

	execute, exists := dispatchers[instruction]

	if !exists {
		help()
		return
	}

	ctx := context.Background()
	arguments := inputs[2:]

	execute(logger, ctx, arguments)
}
