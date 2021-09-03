#!/usr/bin/env bash

set -eo pipefail

SCRIPT_PATH="$(dirname $0)"

list_parsers() {
	go run "$SCRIPT_PATH" --help 2>&1 | grep -oP '// random \K\w*' | sort -u
}

cat <<"EOF" > README.md
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
Licensed under GNU-GPL v3. See [LICENSE](./LICENSE).
EOF
