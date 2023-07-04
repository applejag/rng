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
	"math/big"
	"math/rand"
	"strings"
)

type randomInt struct {
	upper *big.Int
	lower *big.Int
	value *big.Int
}

func (p randomInt) Name() string {
	return "int"
}

func (p randomInt) ParseUpper(value string) (randomUpper, error) {
	var ok bool
	p.upper, ok = big.NewInt(0).SetString(value, 10)
	if !ok {
		return nil, fmt.Errorf("cannot parse as integer: %q", value)
	}
	return p, nil
}

func (p randomInt) Default() randomRange {
	p.upper = big.NewInt(10)
	p.lower = big.NewInt(0)
	return p
}

func (p randomInt) ParseLower(value string) (randomRange, error) {
	var ok bool
	p.lower, ok = big.NewInt(0).SetString(value, 10)
	if !ok {
		return nil, fmt.Errorf("cannot parse as integer: %q", value)
	}
	return p, nil
}

func (p randomInt) DefaultLower() randomRange {
	p.lower = big.NewInt(0)
	return p
}

func (p randomInt) IsLowerLargerThanUpper() bool {
	return p.lower.Cmp(p.upper) > 0
}

func (p randomInt) CalcRandomValue(rnd *rand.Rand) randomValue {
	n := &big.Int{}
	n.Rand(rnd, big.NewInt(0).Add(p.upper, p.lower))
	n.Add(n, p.lower)
	return (*randomIntValue)(n)
}

type randomIntValue big.Int

func (value *randomIntValue) Format(format string) (string, error) {
	bigInt := (*big.Int)(value)
	switch format {
	case "":
		return fmt.Sprint(bigInt), nil
	case "english":
		if bigInt.IsInt64() {
			return englishFormatInt64(bigInt.Int64()), nil
		} else {
			return "", fmt.Errorf("number is too big to be converted into english: %s", bigInt)
		}
	default:
		fmtFormat, hasFormat := intFormats[format]
		if hasFormat {
			return fmt.Sprintf(fmtFormat, bigInt), nil
		}
		return "", errInvalidFormat
	}
}

func (p randomInt) PrintFormatsHelp() {
	fmt.Println(`Formats for int parser:
  --format x               // hexadecimal, ex: c0ffee
  --format X               // hexadecimal, ex: C0FFEE
  --format 0x              // hexadeciaml with prefix, ex: 0xc0ffee
  --format 0X              // hexadeciaml with prefix, ex: 0xC0FFEE
  --format b               // binary, ex: 110000001111111111101110
  --format 0b              // binary with prefix, ex: 0b110000001111111111101110
  --format o               // octal, ex: 60177756
  --format 0o              // octal with prefix, ex: 0o60177756
  --format english         // ex: twelve million six hundred and fourty-eight thousand four hundred and thirty`)
}

var intFormats = map[string]string{
	"x":  "%x\n",
	"0x": "%#x\n",
	"X":  "%X\n",
	"0X": "%#X\n",
	"b":  "%b\n",
	"0b": "%#b\n",
	"o":  "%o\n",
	"0o": "%O\n",
}

var intEnglishFormats = map[int64]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func englishFormatInt64(value int64) string {
	var sb strings.Builder

	if value < 0 {
		sb.WriteString("negative")
		value = -value
	}

	if formatted, ok := intEnglishFormats[value]; ok {
		writeWithSpace(&sb, formatted)
		return sb.String()
	}

	var (
		quintrillions = value / 1e18 // max size of int64 is at quintrillions
		quadrillions  = value % 1e18 / 1e15
		trillions     = value % 1e15 / 1e12
		billions      = value % 1e12 / 1e9
		millions      = value % 1e9 / 1e6
		thousands     = value % 1e6 / 1e3
		mod1000       = value % 1000
	)

	writeIntEnglishFormatsWithSpace(&sb, quintrillions, "quintrillion")
	writeIntEnglishFormatsWithSpace(&sb, quadrillions, "quadrillion")
	writeIntEnglishFormatsWithSpace(&sb, trillions, "trillion")
	writeIntEnglishFormatsWithSpace(&sb, billions, "billion")
	writeIntEnglishFormatsWithSpace(&sb, millions, "million")
	writeIntEnglishFormatsWithSpace(&sb, thousands, "thousand")
	writeIntEnglish100sFormatsWithSpace(&sb, mod1000)

	return sb.String()
}

func writeIntEnglishFormatsWithSpace(sb *strings.Builder, value int64, name string) {
	if value == 0 {
		return
	}
	writeIntEnglish100sFormatsWithSpace(sb, value)
	writeWithSpace(sb, name)
}

func writeIntEnglish100sFormatsWithSpace(sb *strings.Builder, value int64) {
	var (
		hundreds = value % 1000 / 100
		mod100   = value % 100
		tens     = mod100 / 10
		ones     = mod100 % 10
	)
	writeIntEnglishFormatsWithSpace(sb, hundreds, "hundred")
	if mod100 != 0 {
		if mod100 <= 20 || ones == 0 {
			if hundreds != 0 {
				sb.WriteString(" and")
			}
			writeWithSpace(sb, intEnglishFormats[mod100])
		} else {
			if hundreds != 0 {
				sb.WriteString(" and")
			}
			writeWithSpace(sb, intEnglishFormats[tens*10])
			sb.WriteRune('-')
			sb.WriteString(intEnglishFormats[ones])
		}
	}
}

func writeWithSpace(sb *strings.Builder, value string) {
	if sb.Len() > 0 {
		sb.WriteByte(' ')
	}
	sb.WriteString(value)
}
