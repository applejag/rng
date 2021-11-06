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
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"lukechampine.com/uint128"
)

var ErrInvalid = errors.New("uuid: invalid format")

const mask6Bytes uint64 = 0x0000FFFFFFFFFFFF
const mask4Bytes uint64 = 0x00000000FFFFFFFF
const mask2Bytes uint64 = 0x000000000000FFFF

type UUID struct {
	int uint128.Uint128
}

var Min = UUID{}
var Max = UUID{int: uint128.Uint128{Lo: math.MaxUint64, Hi: math.MaxUint64}}

func (uuid UUID) String() string {
	var (
		timeLow          = (uuid.int.Hi >> 32) & mask4Bytes
		timeMid          = (uuid.int.Hi >> 16) & mask2Bytes
		timeHiAndVersion = uuid.int.Hi & mask2Bytes
		clock            = (uuid.int.Lo >> 48) & mask2Bytes
		node             = uuid.int.Lo & mask6Bytes
	)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		timeLow,
		timeMid,
		timeHiAndVersion,
		clock,
		node)
}

func (uuid UUID) GreaterThan(value UUID) bool {
	return uuid.int.Cmp(value.int) > 0
}

func NewRandom() UUID {
	var randBytes [16]byte
	if _, err := rand.Read(randBytes[:]); err != nil {
		panic(err)
	}
	return UUID{int: uint128.FromBytes(randBytes[:])}
}

func NewRandomRange(min, max UUID) UUID {
	diffInt128 := max.int.Sub(min.int)
	fullRandomInt128 := NewRandom().int
	rangedRandomInt128 := fullRandomInt128.Mod(diffInt128)
	randomInt128 := min.int.Add(rangedRandomInt128)
	return UUID{int: randomInt128}
}

func Parse(s string) (UUID, error) {
	if len(s) != 36 {
		return UUID{}, ErrInvalid
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return UUID{}, ErrInvalid
	}
	var err error
	var (
		timeLow          uint64
		timeMid          uint64
		timeHiAndVersion uint64
		clock            uint64
		node             uint64
	)
	if timeLow, err = strconv.ParseUint(s[0:8], 16, 32); err != nil {
		return UUID{}, err
	}
	if timeMid, err = strconv.ParseUint(s[9:13], 16, 16); err != nil {
		return UUID{}, err
	}
	if timeHiAndVersion, err = strconv.ParseUint(s[14:18], 16, 16); err != nil {
		return UUID{}, err
	}
	if clock, err = strconv.ParseUint(s[19:23], 16, 16); err != nil {
		return UUID{}, err
	}
	if node, err = strconv.ParseUint(s[24:36], 16, 48); err != nil {
		return UUID{}, err
	}
	var g = UUID{
		int: uint128.Uint128{
			Hi: timeLow<<32 + timeMid<<16 + timeHiAndVersion,
			Lo: clock<<48 + node,
		},
	}
	return g, nil
}
