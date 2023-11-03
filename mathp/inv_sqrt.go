package mathp

import (
	"math"
)

const (
	_magic32         = 0x5f3759df
	_magic64         = 0x5fe6eb50c7b537a9
	_opf32   float32 = 1.5
	_opf64   float64 = 1.5
	_half32  float32 = 0.5
	_half64  float64 = 0.5
)

func InvSqrt32(num float32) float32 {
	return float32(1.0 / math.Sqrt(float64(num)))
}

func InvSqrt(num float64) float64 {
	return 1.0 / math.Sqrt(num)
}

func FastInvSqrt32(num float32) float32 {
	if num < 0 {
		return float32(math.NaN())
	}

	bits := math.Float32bits(num)
	bits = _magic32 - (bits >> 1)

	res := math.Float32frombits(bits)
	res *= _opf32 - (num * _half32 * res * res)

	return res
}

// FastInvSqrt 平方根倒数速算法.
// Deprecated: 速度提高不大.
func FastInvSqrt(num float64) float64 {
	if num < 0 {
		return math.NaN()
	}

	bits := math.Float64bits(num)
	bits = _magic64 - (bits >> 1)

	res := math.Float64frombits(bits)
	res *= _opf64 - (num * _half64 * res * res)

	return res
}
