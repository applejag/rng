package main

import (
	"fmt"
	"strings"

	"github.com/jilleJr/rng/pkg/uuid"
)

type randomUUID struct {
	upper  uuid.UUID
	lower  uuid.UUID
	format uuidFormat
}

func (p randomUUID) Name() string {
	return "uuid"
}

func (p randomUUID) ParseUpper(value string) (randomUpper, error) {
	var err error
	p.upper, p.format, err = parseUUIDFlexibleFormat(value)
	return p, err
}

func (p randomUUID) Default() randomRange {
	p.upper = uuid.Max
	p.lower = uuid.Min
	p.format = uuidFormatUUID
	return p
}

func (p randomUUID) ParseLower(value string) (randomRange, error) {
	var err error
	p.lower, p.format, err = parseUUIDFlexibleFormat(value)
	return p, err
}

func (p randomUUID) DefaultLower() randomRange {
	p.lower = uuid.Min
	return p
}

func (p randomUUID) IsLowerLargerThanUpper() bool {
	return p.lower.GreaterThan(p.upper)
}

func (p randomUUID) PrintRandomValue() {
	switch p.format {
	case uuidFormatGUID:
		fmt.Printf("{%s}\n", uuid.NewRandomRange(p.lower, p.upper))
	case uuidFormatURN:
		fmt.Printf("urn:uuid:%s\n", uuid.NewRandomRange(p.lower, p.upper))
	default:
		fmt.Println(uuid.NewRandomRange(p.lower, p.upper))
	}
}

type uuidFormat byte

const (
	uuidFormatUUID uuidFormat = iota
	uuidFormatGUID
	uuidFormatURN
)

func parseUUIDFlexibleFormat(value string) (uid uuid.UUID, format uuidFormat, err error) {
	if strings.HasPrefix(value, "urn:uuid:") {
		value = strings.TrimPrefix(value, "urn:uuid:")
		format = uuidFormatURN
	} else if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		value = value[1 : len(value)-1]
		format = uuidFormatGUID
	} else {
		format = uuidFormatUUID
	}

	uid, err = uuid.Parse(value)
	return
}
