package mathp_test

import (
	"testing"

	"github.com/xuender/gpp/mathp"
)

const (
	_32 float32 = 3.33
	_64 float64 = 3.33
)

func BenchmarkInvSqrt32(b *testing.B) {
	for i := 1; i < b.N; i++ {
		mathp.InvSqrt32(_32)
	}
}

func BenchmarkInvSqrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		mathp.InvSqrt(_64)
	}
}

func BenchmarkFastInvSqrt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		mathp.FastInvSqrt(_64)
	}
}

func BenchmarkFastInvSqrt32(b *testing.B) {
	for i := 1; i < b.N; i++ {
		mathp.FastInvSqrt32(_32)
	}
}
