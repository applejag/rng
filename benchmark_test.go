package main

import (
	"testing"
)

func BenchmarkEnglishFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		englishFormatInt64(0xC0FFEE)
	}
}
