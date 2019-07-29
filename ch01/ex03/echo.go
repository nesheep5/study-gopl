package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// EchoA is print args.
func EchoA(out io.Writer, args []string) {
	s, sep := "", ""
	for i := 1; i <= len(args[1:]); i++ {
		s = s + sep + args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// EchoB is print args.
func EchoB(out io.Writer, args []string) {
	fmt.Fprintln(out, strings.Join(args[1:], " "))
}

func main() {
	EchoA(os.Stdout, os.Args)
}
