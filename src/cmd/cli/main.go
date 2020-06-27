package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	os.Exit(run(os.Stdout, os.Args))
}

// run is the main entry point, which is testable. The out should be os.Stdout,
// and in should be os.Args.
func run(out io.Writer, in []string) int {
	fmt.Fprintln(out, "Hello, world!")
	fmt.Fprintln(out, "Got args:", strings.Join(in, ","))
	return 0
}
