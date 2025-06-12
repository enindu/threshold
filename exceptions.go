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

import "errors"

var (
	errNonRoot        error = errors.New("the user must be root")
	errNoInstruction  error = errors.New("the instruction is not found, use \"threshold -h\" or \"threshold --help\" to see help message")
	errInvalidCommand error = errors.New("the command is invalid, use \"threshold -h\" or \"threshold --help\" to see help message")
)
