package main

import (
	"fmt"

	"github.com/jilleJr/rng/pkg/guid"
)

type randomGUID struct {
	upper guid.GUID
	lower guid.GUID
}

func (p randomGUID) Name() string {
	return "guid"
}

func (p randomGUID) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = guid.Parse(value)
	return p, err
}

func (p randomGUID) Default() randomRange {
	p.upper = guid.Max
	p.lower = guid.Min
	return p
}

func (p randomGUID) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = guid.Parse(value)
	return p, err
}

func (p randomGUID) DefaultLower() randomRange {
	p.lower = guid.Min
	return p
}

func (p randomGUID) IsLowerLargerThanUpper() bool {
	return p.lower.GreaterThan(p.upper)
}

func (p randomGUID) PrintRandomValue() {
	fmt.Println(guid.NewRandomRange(p.lower, p.upper))
}
