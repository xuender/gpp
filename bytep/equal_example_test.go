package bytep_test

import (
	"fmt"

	"github.com/xuender/gpp/bytep"
)

func ExampleEqual() {
	fmt.Println(bytep.Equal([]byte("a"), []byte("a")))
	fmt.Println(bytep.Equal([]byte("a"), []byte("b")))

	// Output:
	// true
	// false
}
