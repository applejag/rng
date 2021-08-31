package guid

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"lukechampine.com/uint128"
)

var ErrInvalid = errors.New("guid: invalid format")

const mask6Bytes uint64 = 0x0000FFFFFFFFFFFF
const mask4Bytes uint64 = 0x00000000FFFFFFFF
const mask2Bytes uint64 = 0x000000000000FFFF

type GUID struct {
	int uint128.Uint128
}

var Min = GUID{}
var Max = GUID{int: uint128.Uint128{Lo: math.MaxUint64, Hi: math.MaxUint64}}

func (guid GUID) String() string {
	var (
		timeLow          = (guid.int.Hi >> 32) & mask4Bytes
		timeMid          = (guid.int.Hi >> 16) & mask2Bytes
		timeHiAndVersion = guid.int.Hi & mask2Bytes
		clock            = (guid.int.Lo >> 48) & mask2Bytes
		node             = guid.int.Lo & mask6Bytes
	)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		timeLow,
		timeMid,
		timeHiAndVersion,
		clock,
		node)
}

func (guid GUID) GreaterThan(value GUID) bool {
	return guid.int.Cmp(value.int) > 0
}

func NewRandom() GUID {
	var randBytes [16]byte
	if _, err := rand.Read(randBytes[:]); err != nil {
		panic(err)
	}
	return GUID{int: uint128.FromBytes(randBytes[:])}
}

func NewRandomRange(min, max GUID) GUID {
	diffInt128 := max.int.Sub(min.int)
	fullRandomInt128 := NewRandom().int
	rangedRandomInt128 := fullRandomInt128.Mod(diffInt128)
	randomInt128 := min.int.Add(rangedRandomInt128)
	return GUID{int: randomInt128}
}

func Parse(s string) (GUID, error) {
	if len(s) != 36 {
		return GUID{}, ErrInvalid
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return GUID{}, ErrInvalid
	}
	var err error
	var (
		timeLow          uint64
		timeMid          uint64
		timeHiAndVersion uint64
		clock            uint64
		node             uint64
	)
	if timeLow, err = strconv.ParseUint(s[0:8], 16, 32); err != nil {
		return GUID{}, err
	}
	if timeMid, err = strconv.ParseUint(s[9:13], 16, 16); err != nil {
		return GUID{}, err
	}
	if timeHiAndVersion, err = strconv.ParseUint(s[14:18], 16, 16); err != nil {
		return GUID{}, err
	}
	if clock, err = strconv.ParseUint(s[19:23], 16, 16); err != nil {
		return GUID{}, err
	}
	if node, err = strconv.ParseUint(s[24:36], 16, 48); err != nil {
		return GUID{}, err
	}
	var g = GUID{
		int: uint128.Uint128{
			Hi: timeLow<<32 + timeMid<<16 + timeHiAndVersion,
			Lo: clock<<48 + node,
		},
	}
	return g, nil
}
