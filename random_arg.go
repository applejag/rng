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
	"errors"
	"fmt"
	"math/rand"

	"github.com/spf13/pflag"
)

type randomArg struct{}

var errArgParserInvalidUsage = errors.New("arg parser: cannot parse <min> <max> values")

func (p randomArg) ParseUpper(value string) (randomUpper, error) {
	return nil, errArgParserInvalidUsage
}

func (p randomArg) Default() randomRange {
	return nil
}

func (p randomArg) Name() string {
	return "arg"
}

func (p randomArg) PrintFormatsHelp() {
	fmt.Println(`The arg parser does not support formats.`)
}

func (p randomArg) IsLowerLargerThanUpper() bool {
	return false
}

func (p randomArg) CalcRandomValue(rnd *rand.Rand) randomValue {
	return randomArgValue(pflag.Arg(rnd.Intn(pflag.NArg())))
}

type randomArgValue string

func (value randomArgValue) Format(format string) (string, error) {
	if format != "" {
		return "", errInvalidFormat
	}
	return string(value), nil
}
