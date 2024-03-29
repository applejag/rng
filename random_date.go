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

type randomDate struct {
	upper  time.Time
	lower  time.Time
	layout timeLayout
}

func (p randomDate) Name() string {
	return "date"
}

func (p randomDate) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, p.layout, err = parseDateNicely(value)
	return p, err
}

func (p randomDate) Default() randomRange {
	year := time.Now().Year()
	p.upper = time.Date(year+1, 01, 01, 00, 00, 00, 00, time.UTC)
	p.lower = time.Date(year, 01, 01, 00, 00, 00, 00, time.UTC)
	p.layout = dateLayoutDefault
	return p
}

func (p randomDate) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, p.layout, err = parseDateNicely(value)
	return p, err
}

func (p randomDate) DefaultLower() randomRange {
	p.lower = time.Now()
	return p
}

func (p randomDate) IsLowerLargerThanUpper() bool {
	return p.lower.After(p.upper)
}

func (p randomDate) CalcRandomValue(rnd *rand.Rand) randomValue {
	diffInt := int64(p.upper.Sub(p.lower))
	rndDiff := time.Duration(rnd.Int63n(diffInt))
	rndDate := p.lower.Add(rndDiff)
	return randomDateValue{
		rndDate: rndDate,
		layout:  p.layout,
	}
}

type randomDateValue struct {
	rndDate time.Time
	layout  timeLayout
}

func (value randomDateValue) Format(format string) (string, error) {
	if format == "" {
		return (value.rndDate.Format(value.layout.printLayout)), nil
	} else {
		return (value.rndDate.Format(format)), nil
	}
}

func (p randomDate) PrintFormatsHelp() {
	fmt.Println(`Formats for date parser:
  Any Go time format https://pkg.go.dev/time#pkg-constants
  Such as:
  --format 2006.01.02      // ex: 2021.09.30
  --format 2006            // ex: 2021`)
}

var dateLayoutDefault = timeLayout{"2006-1-2", "2006-01-02"}
var dateLayouts = []timeLayout{
	{"2006-1-2", "2006-01-02"},
	{"6-1-2", "2006-01-02"},
	{"2/1", "02/01"},
	{"2/1-6", "02/01-06"},
	{"2/1-2006", "02/01-2006"},
	{"2 Jan 6", "2 Jan 2006"},
	{"Jan 2", "Jan 2"},
}

func parseDateNicely(value string) (t time.Time, layout timeLayout, err error) {
	for _, layout = range dateLayouts {
		t, err = time.Parse(layout.parseLayout, value)
		if err == nil {
			return
		}
	}
	return
}
