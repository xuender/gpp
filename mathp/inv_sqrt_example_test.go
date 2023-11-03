package mathp_test

import (
	"fmt"

	"github.com/xuender/gpp/mathp"
)

func ExampleInvSqrt() {
	fmt.Println(mathp.InvSqrt(33.3))

	// Output:
	// 0.17329174830334546
}

func ExampleInvSqrt32() {
	fmt.Println(mathp.InvSqrt32(33.3))

	// Output:
	// 0.17329176
}

func ExampleFastInvSqrt() {
	fmt.Println(mathp.FastInvSqrt(33.3))

	// Output:
	// 0.17320113747357305
}

func ExampleFastInvSqrt32() {
	fmt.Println(mathp.FastInvSqrt32(33.3))

	// Output:
	// 0.17320128
}
