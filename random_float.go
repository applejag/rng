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

func (p randomFloat) Name() string {
	return "float"
}

func (p randomFloat) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, err = strconv.ParseFloat(value, 64)
	return p, err
}

func (p randomFloat) Default() randomRange {
	p.upper = 10
	p.lower = 0
	return p
}

func (p randomFloat) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, err = strconv.ParseFloat(value, 64)
	return p, err
}

func (p randomFloat) DefaultLower() randomRange {
	p.lower = 0
	return p
}

func (p randomFloat) IsLowerLargerThanUpper() bool {
	return p.lower > p.upper
}

func (p randomFloat) PrintRandomValue(format string) error {
	var value = lerpFloat64(p.lower, p.upper, rand.Float64())
	if format == "" {
		fmt.Println(value)
		return nil
	}
	fmtFormat, hasFormat := floatFormats[format]
	if hasFormat {
		fmt.Printf(fmtFormat, value)
		return nil
	}
	return errInvalidFormat
}

func (p randomFloat) PrintFormatsHelp() {
	fmt.Println(`Formats for int parser:
  --format e               // scientific notation, ex: 1.23456e+8
  --format E               // scientific notation, ex: 1.23456E+8
  --format f               // decimal, ex: 12345600
  --format x               // hexadeciaml, ex: 0x1.d6f28p+26
  --format X               // hexadeciaml, ex: 0X1.D6F28P+26`)
}

func lerpFloat64(a, b, t float64) float64 {
	return a + (b-a)*t
}

var floatFormats = map[string]string{
	"e": "%e\n",
	"E": "%E\n",
	"f": "%f\n",
	"x": "%x\n",
	"X": "%X\n",
}
