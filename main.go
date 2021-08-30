package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

const (
	version = "v0.1.0"
)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `rnd %s  Copyright (C) 2021  Kalle Jillheden
This program comes with ABSOLUTELY NO WARRANTY; for details type 'rnd --license-w'
This is free software, and you are welcome to redistribute it
under certain conditions; type '--license-c' for details.

Randomizes an integer value by default, but random floats, times, and durations
can be used instead.

rnd                        // random integer [0 - 10)
rnd 15                     // random integer [0 - 15)
rnd 5 15                   // random integer [5 - 15)
rnd 15.0                   // random float [0 - 15)
rnd 1.5e2                  // random float [0 - 15)
rnd 5 1.5e2                // random float [5 - 15)
rnd 15s                    // random duration [0s - 15s)
rnd 5s 15s                 // random duration [5s - 15s)
rnd 10:50                  // random time [now - 10:50)
rnd 10:30 10:15            // random time [15:30 - 10:15]
rnd 2038-01-19             // random date [now - 2038-01-19)
rnd 2021-08-30 2038-01-19  // random date [2021-08-30 - 2038-01-19)

Flags:
`, version)
		pflag.PrintDefaults()
	}

	var flags struct {
		min                   float64
		max                   float64
		useFloat              bool
		showHelp              bool
		showDebug             bool
		showLicenseWarranty   bool
		showLicenseConditions bool
	}

	pflag.BoolVarP(&flags.useFloat, "use-float", "f", false, "enables floating point values (ex: 0.25)")
	pflag.BoolVarP(&flags.showHelp, "help", "h", false, "show this help text")
	pflag.BoolVarP(&flags.showDebug, "debug", "d", false, "show additional info")
	pflag.BoolVarP(&flags.showLicenseConditions, "license-c", "", false, "show license conditions")
	pflag.BoolVarP(&flags.showLicenseWarranty, "license-w", "", false, "show license warranty")
	pflag.Parse()

	if flags.showHelp {
		pflag.Usage()
		os.Exit(0)
	}

	if flags.showLicenseConditions {
		fmt.Println(licenseConditions)
		os.Exit(0)
	}

	if flags.showLicenseWarranty {
		fmt.Println(licenseWarranty)
		os.Exit(0)
	}

	var parser = defaultParser
	var rngRange randomRange

	switch pflag.NArg() {
	case 0:
		rngRange = parser.Default()
	case 1:
		flags.min = 0
		var err error
		flags.max, err = strconv.ParseFloat(pflag.Arg(0), 64)
		if err != nil {
			fmt.Println("err: failed to parse 1st arg <max>:", err)
			os.Exit(2)
		}
	case 2:
		var err error
		flags.min, err = strconv.ParseFloat(pflag.Arg(0), 64)
		if err != nil {
			fmt.Println("err: failed to parse 1st arg <min>:", err)
			os.Exit(2)
		}
		flags.max, err = strconv.ParseFloat(pflag.Arg(1), 64)
		if err != nil {
			fmt.Println("err: failed to parse 2nd arg <max>:", err)
			os.Exit(2)
		}
	}

	if flags.min > flags.max {
		fmt.Println("err: <min> cannot be greater than <max>")
		os.Exit(1)
	}

	if flags.showDebug {
		fmt.Println("min:", flags.min)
		fmt.Println("max:", flags.max)
		fmt.Println("use float:", flags.useFloat)
		fmt.Print("random value: ")
	}

	seedRand()

	if flags.useFloat {
		fmt.Println(lerpFloat64(flags.min, flags.max, rand.Float64()))
	} else {
		var (
			min = int64(flags.min)
			max = int64(flags.max)
		)
		fmt.Println(min + rand.Int63n(max-min))
	}
}

func seedRand() error {
	b := make([]byte, 8)
	if _, err := crand.Read(b); err != nil {
		return fmt.Errorf("failed to create crypto seed: %w", err)
	}
	seed := binary.BigEndian.Uint64(b)
	rand.Seed(int64(seed))
	return nil
}
