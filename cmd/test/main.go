package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	// TODO: test
}

func usage() {
	fmt.Fprintf(os.Stderr, "test\n\n")
	fmt.Fprintf(os.Stderr, "TODO: test.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
