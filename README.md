# rnd - Random number generator CLI

I needed a random number generator. I wrote this super-simple tool.

## Install

Requires Go v1.16+: <https://golang.org/>

```console
$ go install github.com/jilleJr/rng@latest
```

## Usage

```console
$ rng --help
rng v0.2.0  Copyright (C) 2021  Kalle Jillheden
This program comes with ABSOLUTELY NO WARRANTY; for details type 'rng --license-w'
This is free software, and you are welcome to redistribute it
under certain conditions; type '--license-c' for details.

Randomizes an integer value by default, but random floats, times, and durations
can be used instead.

rng                        // random integer [0 - 10)
rng 15                     // random integer [0 - 15)
rng 5 15                   // random integer [5 - 15)
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
  -d, --debug       show additional info
  -h, --help        show this help text
      --license-c   show license conditions
      --license-w   show license warranty
```
