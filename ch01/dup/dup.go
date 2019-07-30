package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//	Dup1(os.Stdin, os.Stdout, os.Args)
	Dup2(os.Stdin, os.Stdout, os.Args)
}

//Dup1 count up duplicate words.
func Dup1(in io.Reader, out io.Writer, args []string) {
	counts := make(map[string]int)
	input := bufio.NewScanner(in)
	fmt.Println("prease input text. and if exit push ctrl-D.")
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%d\t%s\n", n, line)
		}
	}
}

// Dup2 count up duplicate words in files.
func Dup2(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]int)
	filePaths := args[1:]
	if len(filePaths) == 0 {
		fmt.Println("prease input text. and if exit push ctrl-D.")
		countLines(in, counts)
	} else {
		for _, fPath := range filePaths {
			fp, err := os.Open(fPath)
			if err != nil {
				log.Fatal(err)
			}
			countLines(fp, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%d\t%s\n", n, line)
		}
	}
}

func countLines(fp *os.File, counts map[string]int) {
	r := bufio.NewScanner(fp)
	for r.Scan() {
		counts[r.Text()]++
	}
}
