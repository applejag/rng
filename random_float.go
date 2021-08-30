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

func (p randomFloat) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = strconv.ParseFloat(value, 64)
	return p, err
}

func (p randomFloat) Default() randomRange {
	return randomFloat{
		upper: 10,
		lower: 0,
	}
}

func (p randomFloat) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = strconv.ParseFloat(value, 10, 64)
	return p, err
}

func (p randomFloat) DefaultLower() randomRange {
	p.lower = 0
	return p
}

func (p randomFloat) IsLowerLargerThanUpper() bool {
	return p.lower > p.upper
}

func (p randomFloat) PrintRandomValue() {
	fmt.Println(lerpFloat64(p.lower, p.upper, rand.Float64()))
}

func lerpFloat64(a, b, t float64) float64 {
	return a + (b-a)*t
}
