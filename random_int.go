package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type randomInt struct {
	upper int64
	lower int64
}

func (p randomInt) Name() string {
	return "int"
}

func (p randomInt) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = strconv.ParseInt(value, 10, 64)
	return p, err
}

func (p randomInt) Default() randomRange {
	p.upper = 10
	p.lower = 0
	return p
}

func (p randomInt) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = strconv.ParseInt(value, 10, 64)
	return p, err
}

func (p randomInt) DefaultLower() randomRange {
	p.lower = 0
	return p
}

func (p randomInt) IsLowerLargerThanUpper() bool {
	return p.lower > p.upper
}

func (p randomInt) PrintRandomValue() {
	fmt.Println(p.lower + rand.Int63n(p.upper-p.lower))
}
