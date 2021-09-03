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
	version = "v0.5.2"
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
rng 123e4567-e89b-12d3-a456-426614174000 // random UUID [00000000-0000-0000-0000-000000000000 - 123e4567-e89b-12d3-a456-426614174000)
rng -p arg arg1 arg2 ... argN // random arg [arg1, arg2, ..., argN]

Flags:
`, version)
		pflag.PrintDefaults()
	}

	var flags struct {
		parserName            string
		printFormats          []string
		numberOfRandoms       uint
		showHelp              bool
		showHelpParserFormats string
		showDebug             bool
		showLicenseWarranty   bool
		showLicenseConditions bool
	}

	pflag.BoolVarP(&flags.showHelp, "help", "h", false, "show this help text")
	pflag.BoolVarP(&flags.showDebug, "debug", "d", false, "show additional info")
	pflag.BoolVarP(&flags.showLicenseConditions, "license-c", "", false, "show license conditions")
	pflag.BoolVarP(&flags.showLicenseWarranty, "license-w", "", false, "show license warranty")
	pflag.StringVarP(&flags.parserName, "parser", "p", "auto", "force parser")
	pflag.StringSliceVarP(&flags.printFormats, "format", "f", nil, "use custom format(s), separated by comma")
	pflag.StringVarP(&flags.showHelpParserFormats, "help-format", "F", "", "show formats available for parser")
	pflag.UintVarP(&flags.numberOfRandoms, "count", "n", 1, "number of random values to output")
	pflag.Parse()

	if flags.showHelp {
		pflag.Usage()
		os.Exit(0)
	}

	if flags.showHelpParserFormats != "" {
		var rndParser randomParser
		for _, parser := range parsers {
			if parser.Name() == flags.showHelpParserFormats {
				rndParser = parser
				break
			}
		}
		if rndParser == nil {
			fmt.Println(`err: for "-p, --parser" flag: no parser was matched:`, flags.showHelpParserFormats)
			fmt.Println("Available parsers:")
			for _, parser := range parsers {
				fmt.Println("-", parser.Name())
			}
			os.Exit(2)
		}
		rndParser.PrintFormatsHelp()
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
	var parserName = strings.TrimSpace(strings.ToLower(flags.parserName))
	var rndParser randomParser

	switch parserName {
	case "arg":
		if pflag.NArg() == 0 {
			fmt.Println("err: the 'arg' parser needs at least one non-flag argument")
			os.Exit(2)
		}
		rndRange = randomArg{}
		rndParser = randomArg{}
	case "auto":
		switch pflag.NArg() {
		case 0:
			rndRange = defaultRandomRange
			rndParser = defaultParser
		case 1:
			var argUpper = pflag.Arg(0)
			var rndUpper randomUpper
			var err error
			for _, rndParser = range parsers {
				rndUpper, err = rndParser.ParseUpper(argUpper)
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
			for _, rndParser = range parsers {
				rndUpper, err = rndParser.ParseUpper(argUpper)
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
	default:
		for _, parser := range parsers {
			if parser.Name() == parserName {
				rndParser = parser
				break
			}
		}
		if rndParser == nil {
			fmt.Println(`err: for "-p, --parser" flag: no parser was matched:`, flags.parserName)
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
			if err != nil {
				fmt.Println("err: failed to parse <max>:", err)
				os.Exit(2)
			}
			rndRange = rndUpper.DefaultLower()
		case 2:
			var argLower = pflag.Arg(0)
			var argUpper = pflag.Arg(1)
			var err error
			rndUpper, err := rndParser.ParseUpper(argUpper)
			if err != nil {
				fmt.Println("err: failed to parse <max>:", err)
				os.Exit(2)
			}
			rndRange, err = rndUpper.ParseLower(argLower)
			if err != nil {
				fmt.Println("err: failed to parse <min>:", err)
				os.Exit(2)
			}
		}
	}

	if rndRange.IsLowerLargerThanUpper() {
		fmt.Println("err: <min> cannot be greater than <max>")
		os.Exit(1)
	}

	seedRand()

	if len(flags.printFormats) == 0 {
		flags.printFormats = append(flags.printFormats, "")
	}

	for i := uint(0); i < flags.numberOfRandoms; i++ {
		value := rndRange.CalcRandomValue()
		var sb strings.Builder
		for i, format := range flags.printFormats {
			str, err := value.Format(format)
			if err != nil {
				fmt.Println("err:", err)
				rndParser.PrintFormatsHelp()
				os.Exit(1)
			}
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(str)
		}
		fmt.Println(sb.String())
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
