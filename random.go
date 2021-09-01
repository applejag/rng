package main

import (
	"errors"
)

var errMinCannotBeGreaterThanMax = errors.New("<min> cannot be greater than <max>")
var errInvalidFormat = errors.New("invalid parser printing format")

type randomParser interface {
	ParseUpper(value string) (randomUpper, error)
	Default() randomRange
	Name() string
}

type randomUpper interface {
	ParseLower(value string) (randomRange, error)
	DefaultLower() randomRange
}

type randomRange interface {
	IsLowerLargerThanUpper() bool
	PrintRandomValue(format string) error
	PrintFormatsHelp()
}

var parsers = []randomParser{
	randomInt{},
	randomFloat{},
	randomDuration{},
	randomTime{},
	randomDate{},
	randomUUID{},
}

var defaultParser randomParser = randomInt{}
var defaultRandomRange randomRange = randomInt{
	upper: 10,
	lower: 0,
}
