package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	Dup1(os.Stdin, os.Stdout, os.Args)
}

//Dup1 count up duplicate words.
func Dup1(in io.Reader, out io.Writer, args []string) {
	counts := make(map[string]int)
	input := bufio.NewScanner(in)
	for input.Scan() {
		line := input.Text()
		if len(line) < 1 {
			break
		}
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%d\t%s\n", n, line)
		}
	}
}
