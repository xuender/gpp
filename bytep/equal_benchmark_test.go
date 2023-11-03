package bytep_test

import (
	"bytes"
	"testing"

	"github.com/xuender/gpp/bytep"
)

// nolint
var (
	_dataA = bytes.Repeat([]byte("test data A"), 1000_000)
	_dataB = bytes.Repeat([]byte("test data B"), 1000_001)
)

func BenchmarkEqual_bytep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytep.Equal(_dataA, _dataB)
	}
}

func BenchmarkEqual_bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Equal(_dataA, _dataB)
	}
}
