<!--
SPDX-FileCopyrightText: 2021 Kalle Jillheden

SPDX-License-Identifier: CC0-1.0
-->

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
rng v0.5.2  Copyright (C) 2021  Kalle Jillheden
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

## Sample use-cases

### Bash/Zsh

```bash
# Opens a random episode from musicforprogramming.net
alias musicforprogramming='xdg-open "https://musicforprogramming.net/?$(rng 1 63 -f english | tr -d -)"'
```

### PowerShell

```powershell
# Opens a random episode from musicforprogramming.net
function musicforprogramming { start "https://musicforprogramming.net/?$((rng 1 63 -f english).Replace('-',''))" }
```

## Formats

### Parser: UUID

```console
$ rng --help-format UUID
Formats for UUID parser:
  --format uuid            // UUID, ex:
                           //  123e4567-e89b-12d3-a456-426614174000
  --format UUID            // UUID, ex:
                           //  123E4567-E89B-12D3-A456-426614174000

  --format urn             // Uniform Resource Name, ex:
                           //  urn:uuid:123e4567-e89b-12d3-a456-426614174000
  --format URN             // Uniform Resource Name, ex:
                           //  URN:UUID:123E4567-E89B-12D3-A456-426614174000

  --format GUID            // Microsoft GUID, ex:
                           //  {123E4567-E89B-12D3-A456-426614174000}
  --format guid            // Microsoft GUID, ex:
                           //  {123e4567-e89b-12d3-a456-426614174000}

```

### Parser: arg

```console
$ rng --help-format arg
The arg parser does not support formats.

```

### Parser: date

```console
$ rng --help-format date
Formats for date parser:
  Any Go time format https://pkg.go.dev/time#pkg-constants
  Such as:
  --format 2006.01.02      // ex: 2021.09.30
  --format 2006            // ex: 2021

```

### Parser: duration

```console
$ rng --help-format duration
Formats for duration parser:
  --format ns              // nanoseconds, ex: 86400000000000
  --format us              // microseconds, ex: 86400000000
  --format ms              // milliseconds, ex: 86400000
  --format s               // seconds, ex: 86400
  --format m               // minutes, ex: 1440
  --format h               // hours, ex: 24

```

### Parser: float

```console
$ rng --help-format float
Formats for int parser:
  --format e               // scientific notation, ex: 1.23456e+8
  --format E               // scientific notation, ex: 1.23456E+8
  --format f               // decimal, ex: 12345600
  --format x               // hexadeciaml, ex: 0x1.d6f28p+26
  --format X               // hexadeciaml, ex: 0X1.D6F28P+26

```

### Parser: int

```console
$ rng --help-format int
Formats for int parser:
  --format x               // hexadecimal, ex: c0ffee
  --format X               // hexadecimal, ex: C0FFEE
  --format 0x              // hexadeciaml with prefix, ex: 0xc0ffee
  --format 0X              // hexadeciaml with prefix, ex: 0xC0FFEE
  --format b               // binary, ex: 110000001111111111101110
  --format 0b              // binary with prefix, ex: 0b110000001111111111101110
  --format o               // octal, ex: 60177756
  --format 0o              // octal with prefix, ex: 0o60177756
  --format english         // ex: twelve million six hundred and fourty-eight thousand four hundred and thirty

```

### Parser: time

```console
$ rng --help-format time
Formats for time parser:
  Any Go time format https://pkg.go.dev/time#pkg-constants
  Such as:
  --format 15h04m          // ex: 19h33m
  --format 3PM             // ex: 7PM

```

## License

Written and maintained by [@jilleJr](https://github.com/jilleJr).
Licensed under the GNU GPL 3.0 or later, or the CC0 1.0, depending on the file.

This repository is [REUSE](https://reuse.software/) compliant.
