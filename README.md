# rnd - Random number generator CLI

I needed a random number generator. I wrote this super-simple tool.

## Usage

```console
$ rnd --help
rnd v0.2.0  Copyright (C) 2021  Kalle Jillheden
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
  -d, --debug       show additional info
  -h, --help        show this help text
      --license-c   show license conditions
      --license-w   show license warranty
```
