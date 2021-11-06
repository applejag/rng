// SPDX-FileCopyrightText: 2021 Kalle Jillheden
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package main

import "strconv"

func parseUint32(value string, base int) (uint32, error) {
	var result, err = strconv.ParseUint(value, base, 32)
	if err != nil {
		return 0, err
	}
	return uint32(result), nil
}

func parseUint16(value string, base int) (uint16, error) {
	var result, err = strconv.ParseUint(value, base, 16)
	if err != nil {
		return 0, err
	}
	return uint16(result), nil
}
