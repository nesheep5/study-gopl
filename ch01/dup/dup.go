package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//	Dup1(os.Stdin, os.Stdout, os.Args)
	//	Dup2(os.Stdin, os.Stdout, os.Args)
	Dup3(os.Stdin, os.Stdout, os.Args)

}

//Dup1 count up duplicate words.
func Dup1(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]int)
	fmt.Println("please input text. and when exit, press <Ctrl-D>.")
	countLines(in, counts)
	printDupCounts(out, counts)
}

// Dup2 count up duplicate words in files.(Stream Mode)
func Dup2(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]int)
	filePaths := args[1:]
	if len(filePaths) == 0 {
		fmt.Println("please input text. and when exit, press <Ctrl-D>.")
		countLines(in, counts)
	} else {
		for _, fPath := range filePaths {
			fp, err := os.Open(fPath)
			if err != nil {
				log.Fatal(err)
			}
			countLines(fp, counts)
			fp.Close()
		}
	}
	printDupCounts(out, counts)
}

// Dup3 count up duplicate words in files.(All Read Mode)
func Dup3(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]int)
	filePaths := args[1:]
	for _, fPath := range filePaths {
		data, err := ioutil.ReadFile(fPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	printDupCounts(out, counts)
}

func countLines(fp *os.File, counts map[string]int) {
	s := bufio.NewScanner(fp)
	for s.Scan() {
		counts[s.Text()]++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func printDupCounts(out io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%d\t%s\n", n, line)
		}
	}
}
