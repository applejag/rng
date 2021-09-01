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

func (p randomArg) PrintRandomValue() {
	fmt.Println(pflag.Arg(rand.Intn(pflag.NArg())))
}
