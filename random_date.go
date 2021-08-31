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

func (p randomDate) PrintRandomValue() {
	diffInt := int64(p.upper.Sub(p.lower))
	rndDiff := time.Duration(rand.Int63n(diffInt))
	fmt.Println(p.lower.Add(rndDiff).Format(p.layout.printLayout))
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
