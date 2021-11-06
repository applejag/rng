#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2021 Kalle Jillheden
#
# SPDX-License-Identifier: CC0-1.0

set -eo pipefail

SCRIPT_PATH="$(dirname $0)"

list_parsers() {
	go run "$SCRIPT_PATH" --help 2>&1 | grep -o '// random \w*' | awk '{print $3}' | sort -u
}

cat <<"EOF" > README.md
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
EOF

go run "$SCRIPT_PATH" --help >> README.md 2>> README.md

cat <<"EOF" >> README.md
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

EOF

for PARSER in $(list_parsers); do
	echo "### Parser: $PARSER" >> README.md
	echo '' >> README.md
	echo '```console' >> README.md
	echo "\$ rng --help-format $PARSER" >> README.md
	go run "$SCRIPT_PATH" --help-format "$PARSER" >> README.md 2>> README.md
	echo '' >> README.md
	echo '```' >> README.md
	echo '' >> README.md
done

cat <<"EOF" >> README.md
## License

Written and maintained by [@jilleJr](https://github.com/jilleJr).
Licensed under the GNU GPL 3.0 or later, or the CC0 1.0, depending on the file.

This repository is [REUSE](https://reuse.software/) compliant.
EOF
