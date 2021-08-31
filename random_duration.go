package main

import (
	"fmt"
	"math/rand"
	"time"
)

type randomDuration struct {
	upper time.Duration
	lower time.Duration
}

func (p randomDuration) Name() string {
	return "duration"
}

func (p randomDuration) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = time.ParseDuration(value)
	return p, err
}

func (p randomDuration) Default() randomRange {
	p.lower = 0
	p.upper = 10 * time.Second
	return p
}

func (p randomDuration) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = time.ParseDuration(value)
	return p, err
}

func (p randomDuration) DefaultLower() randomRange {
	p.lower = 0
	return p
}

func (p randomDuration) IsLowerLargerThanUpper() bool {
	return p.lower > p.upper
}

func (p randomDuration) PrintRandomValue() {
	upperInt := int64(p.upper)
	lowerInt := int64(p.lower)
	fmt.Println(p.lower + time.Duration(rand.Int63n(upperInt-lowerInt)))
}
