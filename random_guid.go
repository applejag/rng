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

func (p randomUUID) CalcRandomValue() randomValue {
	var value = uuid.NewRandomRange(p.lower, p.upper)
	return randomUUIDValue{
		uuid:   value,
		format: p.format,
	}
}

type randomUUIDValue struct {
	uuid   uuid.UUID
	format uuidFormat
}

func (value randomUUIDValue) PrintRandomValue(format string) (string, error) {
	switch format {
	case "":
		switch value.format {
		case uuidFormatGUID:
			return fmt.Sprintf("{%s}\n", value.uuid), nil
		case uuidFormatURN:
			return fmt.Sprintf("urn:uuid:%s\n", value.uuid), nil
		default:
			return fmt.Sprintln(value.uuid), nil
		}
	case "uuid":
		return fmt.Sprintln(value.uuid), nil
	case "UUID":
		return fmt.Sprintln(strings.ToUpper(value.uuid.String())), nil
	case "urn":
		return fmt.Sprintf("urn:uuid:%s\n", value.uuid), nil
	case "URN":
		return fmt.Sprintf("URN:UUID:%s\n", strings.ToUpper(value.uuid.String())), nil
	case "guid":
		return fmt.Sprintf("{%s}\n", value.uuid), nil
	case "GUID":
		return fmt.Sprintf("{%s}\n", strings.ToUpper(value.uuid.String())), nil
	default:
		return "", errInvalidFormat
	}
}

func (p randomUUID) PrintFormatsHelp() {
	fmt.Println(`Formats for UUID parser:
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
                           //  {123e4567-e89b-12d3-a456-426614174000}`)
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
