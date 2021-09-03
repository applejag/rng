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

func (p randomDuration) CalcRandomValue() randomValue {
	upperInt := int64(p.upper)
	lowerInt := int64(p.lower)
	rndDuration := p.lower + time.Duration(rand.Int63n(upperInt-lowerInt))
	return randomDurationValue(rndDuration)
}

type randomDurationValue time.Duration

func (value randomDurationValue) Format(format string) (string, error) {
	durValue := time.Duration(value)
	switch format {
	case "":
		return fmt.Sprintln(durValue), nil
	case "ns":
		return fmt.Sprintln(int64(durValue)), nil
	case "us":
		const divNanoToMicro = 1000
		return fmt.Sprintln(int64(durValue) / divNanoToMicro), nil
	case "ms":
		const divNanoToMilli = 1000 * 1000
		return fmt.Sprintln(int64(durValue) / divNanoToMilli), nil
	case "s":
		const divNanoToSecond = 1000 * 1000 * 1000
		return fmt.Sprintln(int64(durValue) / divNanoToSecond), nil
	case "m":
		const divNanoToMinute = 1000 * 1000 * 1000 * 60
		return fmt.Sprintln(int64(durValue) / divNanoToMinute), nil
	case "h":
		const divNanoToHour = 1000 * 1000 * 1000 * 60 * 60
		return fmt.Sprintln(int64(durValue) / divNanoToHour), nil
	default:
		return "", errInvalidFormat
	}
}

func (p randomDuration) PrintFormatsHelp() {
	fmt.Println(`Formats for duration parser:
  --format ns              // nanoseconds, ex: 86400000000000
  --format us              // microseconds, ex: 86400000000
  --format ms              // milliseconds, ex: 86400000
  --format s               // seconds, ex: 86400
  --format m               // minutes, ex: 1440
  --format h               // hours, ex: 24`)
}
