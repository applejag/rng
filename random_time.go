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

func (p randomTime) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, p.layout, err = parseTimeNicely(value)
	return p, err
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

func (p randomTime) PrintRandomValue() {
	diffInt := int64(p.upper.Sub(p.lower))
	rndDiff := time.Duration(rand.Int63n(diffInt))
	fmt.Println(p.lower.Add(rndDiff).Format(p.layout.printLayout))
}

var timeLayouts = []timeLayout{
	{"15:4", "15:04"},
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