package main

import (
	"fmt"
	"math/rand"

	"github.com/spf13/pflag"
)

type randomArg struct{}

func (p randomArg) IsLowerLargerThanUpper() bool {
	return false
}

func (p randomArg) PrintRandomValue(format string) error {
	if format != "" {
		return errInvalidFormat
	}
	var value = pflag.Arg(rand.Intn(pflag.NArg()))
	fmt.Println(value)
	return nil
}

func (p randomArg) PrintFormatsHelp() {
	fmt.Println(`The arg parser does not support formats.`)
}
