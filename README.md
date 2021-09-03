# rng - Random number generator CLI

I needed a random number generator. I wrote this super-simple tool.

## Install

Requires Go v1.16+: <https://golang.org/>

```console
$ go install github.com/jilleJr/rng@latest
```

## Usage

```console
$ rng --help
rng v0.5.0  Copyright (C) 2021  Kalle Jillheden
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
  -n, --count uint           number of random values to output (default 1)
  -d, --debug                show additional info
  -f, --format strings       use custom format(s), separated by comma
  -h, --help                 show this help text
  -F, --help-format string   show formats available for parser
      --license-c            show license conditions
      --license-w            show license warranty
  -p, --parser string        force parser (default "auto")
```
