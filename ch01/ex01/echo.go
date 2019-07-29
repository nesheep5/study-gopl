package main

import (
	"fmt"
	"io"
	"os"
)

// Echo is print args.
func Echo(out io.Writer, args []string) {
	s, sep := "", ""
	for i := 1; i <= len(args[1:]); i++ {
		s = s + sep + args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func main() {
	Echo(os.Stdout, os.Args)
}
