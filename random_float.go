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

import (
	"fmt"
	"math/rand"
	"strconv"
)

type randomFloat struct {
	upper float64
	lower float64
}

func (p randomFloat) Name() string {
	return "float"
}

func (p randomFloat) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = strconv.ParseFloat(value, 64)
	return p, err
}

func (p randomFloat) Default() randomRange {
	p.upper = 10
	p.lower = 0
	return p
}

func (p randomFloat) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = strconv.ParseFloat(value, 64)
	return p, err
}

func (p randomFloat) DefaultLower() randomRange {
	p.lower = 0
	return p
}

func (p randomFloat) IsLowerLargerThanUpper() bool {
	return p.lower > p.upper
}

type randomFloatValue float64

func (p randomFloat) CalcRandomValue(rnd *rand.Rand) randomValue {
	return randomFloatValue(lerpFloat64(p.lower, p.upper, rnd.Float64()))
}

func (value randomFloatValue) Format(format string) (string, error) {
	if format == "" {
		return fmt.Sprint(value), nil
	}
	fmtFormat, hasFormat := floatFormats[format]
	if hasFormat {
		return fmt.Sprintf(fmtFormat, value), nil
	}
	return "", errInvalidFormat
}

func (p randomFloat) PrintFormatsHelp() {
	fmt.Println(`Formats for int parser:
  --format e               // scientific notation, ex: 1.23456e+8
  --format E               // scientific notation, ex: 1.23456E+8
  --format f               // decimal, ex: 12345600
  --format x               // hexadeciaml, ex: 0x1.d6f28p+26
  --format X               // hexadeciaml, ex: 0X1.D6F28P+26`)
}

func lerpFloat64(a, b, t float64) float64 {
	return a + (b-a)*t
}

var floatFormats = map[string]string{
	"e": "%e\n",
	"E": "%E\n",
	"f": "%f\n",
	"x": "%x\n",
	"X": "%X\n",
}
