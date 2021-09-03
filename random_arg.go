package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/spf13/pflag"
)

type randomArg struct{}

var errArgParserInvalidUsage = errors.New("arg parser: cannot parse <min> <max> values")

func (p randomArg) ParseUpper(value string) (randomUpper, error) {
	return nil, errArgParserInvalidUsage
}

func (p randomArg) Default() randomRange {
	return nil
}

func (p randomArg) Name() string {
	return "arg"
}

func (p randomArg) PrintFormatsHelp() {
	fmt.Println(`The arg parser does not support formats.`)
}

func (p randomArg) IsLowerLargerThanUpper() bool {
	return false
}

func (p randomArg) CalcRandomValue() randomValue {
	return randomArgValue(pflag.Arg(rand.Intn(pflag.NArg())))
}

type randomArgValue string

func (value randomArgValue) Format(format string) (string, error) {
	if format != "" {
		return "", errInvalidFormat
	}
	return string(value), nil
}
