package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Echo is print args.
func Echo(out io.Writer, args []string) {
	fmt.Fprintln(out, strings.Join(args[1:], " "))
}

func main() {
	Echo(os.Stdout, os.Args)
}
