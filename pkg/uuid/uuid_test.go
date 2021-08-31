package uuid

import (
	"testing"

	"lukechampine.com/uint128"
)

func TestParse(t *testing.T) {
	const wantHi uint64 = 0x123456789abcdef1
	const wantLo uint64 = 0x23456789abcdef12
	got, err := Parse("12345678-9abc-def1-2345-6789abcdef12")
	if err != nil {
		t.Fatal(err)
	}
	if got.int.Hi != wantHi {
		t.Errorf("want hi: %#016x, got: %#016x", wantHi, got.int.Hi)
	}
	if got.int.Lo != wantLo {
		t.Errorf("want lo: %#016x, got: %#016x", wantLo, got.int.Lo)
	}
}

func TestString(t *testing.T) {
	const hi uint64 = 0x123456789abcdef1
	const lo uint64 = 0x23456789abcdef12
	const want = "12345678-9abc-def1-2345-6789abcdef12"
	guid := UUID{int: uint128.Uint128{Hi: hi, Lo: lo}}
	got := guid.String()
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
