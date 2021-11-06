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

package uuid

import (
	"testing"

	"lukechampine.com/uint128"
)

func TestParse(t *testing.T) {
	const wantHi uint64 = 0x123456789abcdef1
	const wantLo uint64 = 0x23456789abcdef12
	got, err := Parse("12345678-9abc-def1-2345-6789abcdef12")
	if err != nil {
		t.Fatal(err)
	}
	if got.int.Hi != wantHi {
		t.Errorf("want hi: %#016x, got: %#016x", wantHi, got.int.Hi)
	}
	if got.int.Lo != wantLo {
		t.Errorf("want lo: %#016x, got: %#016x", wantLo, got.int.Lo)
	}
}

func TestString(t *testing.T) {
	const hi uint64 = 0x123456789abcdef1
	const lo uint64 = 0x23456789abcdef12
	const want = "12345678-9abc-def1-2345-6789abcdef12"
	value := UUID{int: uint128.Uint128{Hi: hi, Lo: lo}}
	got := value.String()
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
