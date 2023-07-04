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
	"math/rand"
)

var errMinCannotBeGreaterThanMax = errors.New("<min> cannot be greater than <max>")
var errInvalidFormat = errors.New("invalid parser printing format")

type randomParser interface {
	ParseUpper(value string) (randomUpper, error)
	Default() randomRange
	Name() string
	PrintFormatsHelp()
}

type randomUpper interface {
	ParseLower(value string) (randomRange, error)
	DefaultLower() randomRange
}

type randomRange interface {
	IsLowerLargerThanUpper() bool
	CalcRandomValue(rnd *rand.Rand) randomValue
}

type randomValue interface {
	Format(format string) (string, error)
}

var parsers = []randomParser{
	randomArg{},
	randomInt{},
	randomFloat{},
	randomDuration{},
	randomTime{},
	randomDate{},
	randomUUID{},
}

var defaultParser randomParser = randomInt{}
var defaultRandomRange randomRange = randomInt{}.Default()
