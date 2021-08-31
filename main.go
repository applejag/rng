package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const (
	version = "v0.3.0-rc"
)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `rng %s  Copyright (C) 2021  Kalle Jillheden
This program comes with ABSOLUTELY NO WARRANTY; for details type 'rng --license-w'
This is free software, and you are welcome to redistribute it
under certain conditions; type '--license-c' for details.

Randomizes an integer value by default, but random floats, times, and durations
can be used instead.

rng                        // random int [0 - 10)
rng 15                     // random int [0 - 15)
rng 5 15                   // random int [5 - 15)
rng 15.0                   // random float [0 - 15)
rng 1.5e2                  // random float [0 - 15)
rng 5 1.5e2                // random float [5 - 15)
rng 15s                    // random duration [0s - 15s)
rng 5s 15s                 // random duration [5s - 15s)
rng 10:50                  // random time [now - 10:50)
rng 10:30 10:15            // random time [15:30 - 10:15]
rng 2038-01-19             // random date [now - 2038-01-19)
rng 2021-08-30 2038-01-19  // random date [2021-08-30 - 2038-01-19)

Flags:
`, version)
		pflag.PrintDefaults()
	}

	var flags struct {
		parserName            string
		showHelp              bool
		showDebug             bool
		showLicenseWarranty   bool
		showLicenseConditions bool
	}

	pflag.BoolVarP(&flags.showHelp, "help", "h", false, "show this help text")
	pflag.BoolVarP(&flags.showDebug, "debug", "d", false, "show additional info")
	pflag.BoolVarP(&flags.showLicenseConditions, "license-c", "", false, "show license conditions")
	pflag.BoolVarP(&flags.showLicenseWarranty, "license-w", "", false, "show license warranty")
	pflag.StringVarP(&flags.parserName, "parser", "p", "auto", "force parser")
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

	var rndRange randomRange

	if flags.parserName != "auto" {
		var randomType = strings.TrimSpace(strings.ToLower(flags.parserName))
		var rndParser randomParser
		for _, parser := range parsers {
			if parser.Name() == randomType {
				rndParser = parser
				break
			}
		}
		if rndParser == nil {
			fmt.Println(`err: for "-p, --parser" flag: no parser was matched:`, randomType)
			fmt.Println("Available parsers:")
			for _, parser := range parsers {
				fmt.Println("-", parser.Name())
			}
			os.Exit(2)
		}
		switch pflag.NArg() {
		case 0:
			rndRange = rndParser.Default()
		case 1:
			var argUpper = pflag.Arg(0)
			var rndUpper randomUpper
			var err error
			rndUpper, err = rndParser.ParseUpper(argUpper)
			if err == nil {
				fmt.Println("err: failed to parse <max>:", err)
				os.Exit(2)
			}
			rndRange = rndUpper.DefaultLower()
		case 2:
			var argLower = pflag.Arg(0)
			var argUpper = pflag.Arg(1)
			var err error
			rndUpper, err := rndParser.ParseUpper(argUpper)
			if err == nil {
				fmt.Println("err: failed to parse <max>:", err)
				os.Exit(2)
			}
			rndRange, err = rndUpper.ParseLower(argLower)
			if err == nil {
				fmt.Println("err: failed to parse <min>:", err)
				os.Exit(2)
			}
		}
	} else {
		switch pflag.NArg() {
		case 0:
			rndRange = defaultRandomRange
		case 1:
			var argUpper = pflag.Arg(0)
			var rndUpper randomUpper
			var err error
			for _, p := range parsers {
				rndUpper, err = p.ParseUpper(argUpper)
				if err == nil {
					break
				}
			}
			if err != nil || rndUpper == nil {
				fmt.Println("err: failed to find matching format")
				os.Exit(2)
			}
			rndRange = rndUpper.DefaultLower()
		case 2:
			var argLower = pflag.Arg(0)
			var argUpper = pflag.Arg(1)
			var rndUpper randomUpper
			var err error
			for _, p := range parsers {
				rndUpper, err = p.ParseUpper(argUpper)
				if err != nil {
					continue
				}
				rndRange, err = rndUpper.ParseLower(argLower)
				if err == nil {
					break
				}
			}
			if err != nil || rndRange == nil {
				fmt.Println("err: failed to find matching format")
				os.Exit(2)
			}
		}
	}

	if rndRange.IsLowerLargerThanUpper() {
		fmt.Println("err: <min> cannot be greater than <max>")
		os.Exit(1)
	}

	seedRand()

	rndRange.PrintRandomValue()
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
