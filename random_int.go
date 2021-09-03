package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type randomInt struct {
	upper int64
	lower int64
	value int64
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

func (p randomInt) CalcRandomValue() randomValue {
	return randomIntValue(p.lower + rand.Int63n(p.upper-p.lower))
}

type randomIntValue int64

func (value randomIntValue) PrintRandomValue(format string) (string, error) {
	switch format {
	case "":
		return fmt.Sprint(value), nil
	case "english":
		return englishFormatInt64(int64(value)), nil
	default:
		fmtFormat, hasFormat := intFormats[format]
		if hasFormat {
			return fmt.Sprintf(fmtFormat, value), nil
		}
		return "", errInvalidFormat
	}
}

func (p randomIntValue) PrintFormatsHelp() {
	fmt.Println(`Formats for int parser:
  --format x               // hexadecimal, ex: c0ffee
  --format X               // hexadecimal, ex: C0FFEE
  --format 0x              // hexadeciaml with prefix, ex: 0xc0ffee
  --format 0X              // hexadeciaml with prefix, ex: 0xC0FFEE
  --format b               // binary, ex: 110000001111111111101110
  --format 0b              // binary with prefix, ex: 0b110000001111111111101110
  --format o               // octal, ex: 60177756
  --format 0o              // octal with prefix, ex: 0o60177756
  --format english         // ex: twelve million six hundred and fourty-eight thousand four hundred and thirty`)
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

var intEnglishFormats = map[int64]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func englishFormatInt64(value int64) string {
	if formatted, ok := intEnglishFormats[value]; ok {
		return formatted
	}

	var (
		sb            strings.Builder
		quintrillions = value / 1e18 // max size of int64 is at quintrillions
		quadrillions  = value % 1e18 / 1e15
		trillions     = value % 1e15 / 1e12
		billions      = value % 1e12 / 1e9
		millions      = value % 1e9 / 1e6
		thousands     = value % 1e6 / 1e3
		mod1000       = value % 1000
	)

	writeIntEnglishFormatsWithSpace(&sb, quintrillions, "quintrillion")
	writeIntEnglishFormatsWithSpace(&sb, quadrillions, "quadrillion")
	writeIntEnglishFormatsWithSpace(&sb, trillions, "trillion")
	writeIntEnglishFormatsWithSpace(&sb, billions, "billion")
	writeIntEnglishFormatsWithSpace(&sb, millions, "million")
	writeIntEnglishFormatsWithSpace(&sb, thousands, "thousand")
	writeIntEnglish100sFormatsWithSpace(&sb, mod1000)

	return sb.String()
}

func writeIntEnglishFormatsWithSpace(sb *strings.Builder, value int64, name string) {
	if value == 0 {
		return
	}
	writeIntEnglish100sFormatsWithSpace(sb, value)
	writeWithSpace(sb, name)
}

func writeIntEnglish100sFormatsWithSpace(sb *strings.Builder, value int64) {
	var (
		hundreds = value % 1000 / 100
		mod100   = value % 100
		tens     = mod100 / 10
		ones     = mod100 % 10
	)
	writeIntEnglishFormatsWithSpace(sb, hundreds, "hundred")
	if mod100 != 0 {
		if mod100 <= 20 || ones == 0 {
			if hundreds != 0 {
				sb.WriteString(" and")
			}
			writeWithSpace(sb, intEnglishFormats[mod100])
		} else {
			if hundreds != 0 {
				sb.WriteString(" and")
			}
			writeWithSpace(sb, intEnglishFormats[tens*10])
			sb.WriteRune('-')
			sb.WriteString(intEnglishFormats[ones])
		}
	}
}

func writeWithSpace(sb *strings.Builder, value string) {
	if sb.Len() > 0 {
		sb.WriteByte(' ')
	}
	sb.WriteString(value)
}
