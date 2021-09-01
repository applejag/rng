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

func (p randomInt) PrintRandomValue(format string) error {
	var value = p.lower + rand.Int63n(p.upper-p.lower)
	if format == "" {
		fmt.Println(value)
		return nil
	}
	fmtFormat, hasFormat := intFormats[format]
	if hasFormat {
		fmt.Printf(fmtFormat, value)
		return nil
	}
	return errInvalidFormat
}

func (p randomInt) PrintFormatsHelp() {
	fmt.Println(`Formats for int parser:
  --format x               // hexadecimal, ex: c0ffee
  --format X               // hexadecimal, ex: C0FFEE
  --format 0x              // hexadeciaml with prefix, ex: 0xc0ffee
  --format 0X              // hexadeciaml with prefix, ex: 0xC0FFEE
  --format b               // binary, ex: 110000001111111111101110
  --format 0b              // binary with prefix, ex: 0b110000001111111111101110
  --format o               // octal, ex: 60177756
  --format 0o              // octal with prefix, ex: 0o60177756`)
}

var intFormats = map[string]string{
	"x":  "%x\n",
	"0x": "%#x\n",
	"X":  "%X\n",
	"0X": "%#X\n",
	"b":  "%b\n",
	"0b": "%#b\n",
	"o":  "%o\n",
	"0o": "%O\n",
}
