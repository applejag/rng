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
	"time"
)

type timeLayout struct {
	parseLayout string
	printLayout string
}

type randomTime struct {
	upper  time.Time
	lower  time.Time
	layout timeLayout
}

func (p randomTime) Name() string {
	return "time"
}

func (p randomTime) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, p.layout, err = parseTimeNicely(value)
	return p, err
}

func (p randomTime) Default() randomRange {
	p.upper = time.Date(2006, 01, 03, 00, 00, 00, 00, time.UTC)
	p.lower = time.Date(2006, 01, 02, 00, 00, 00, 00, time.UTC)
	p.layout = timeLayoutDefault
	return p
}

func (p randomTime) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, p.layout, err = parseTimeNicely(value)
	return p, err
}

func (p randomTime) DefaultLower() randomRange {
	p.lower = time.Now()
	return p
}

func (p randomTime) IsLowerLargerThanUpper() bool {
	return p.lower.After(p.upper)
}

func (p randomTime) CalcRandomValue(rnd *rand.Rand) randomValue {
	diffInt := int64(p.upper.Sub(p.lower))
	rndDiff := time.Duration(rnd.Int63n(diffInt))
	rndTime := p.lower.Add(rndDiff)
	return randomTimeValue{
		rndTime: rndTime,
		layout:  p.layout,
	}
}

type randomTimeValue struct {
	rndTime time.Time
	layout  timeLayout
}

func (value randomTimeValue) Format(format string) (string, error) {
	if format == "" {
		return (value.rndTime.Format(value.layout.printLayout)), nil
	} else {
		return (value.rndTime.Format(format)), nil
	}
}

func (p randomTime) PrintFormatsHelp() {
	fmt.Println(`Formats for time parser:
  Any Go time format https://pkg.go.dev/time#pkg-constants
  Such as:
  --format 15h04m          // ex: 19h33m
  --format 3PM             // ex: 7PM`)
}

var timeLayoutDefault = timeLayout{"15:4", "15:04"}

var timeLayouts = []timeLayout{
	timeLayoutDefault,
	{"15:4:5", "15:04:05"},
	{"15:4:5.999999999", "15:04:05.999999999"},
	{"3:4PM", "3:04PM"},
	{"3:4:5PM", "3:04:05PM"},
	{"3:4:5.999999999PM", "3:04:05.999999999PM"},
}

func parseTimeNicely(value string) (t time.Time, layout timeLayout, err error) {
	for _, layout = range timeLayouts {
		t, err = time.Parse(layout.parseLayout, value)
		if err == nil {
			return
		}
	}
	return
}
