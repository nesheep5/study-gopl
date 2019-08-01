package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	Dup2(os.Stdin, os.Stdout, os.Args)

}

type fileNameMap map[string]int

// Dup2 count up and display filename to duplicate words in files.(Stream Mode)
func Dup2(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]fileNameMap)
	filePaths := args[1:]
	if len(filePaths) == 0 {
		fmt.Println("no files.")
		return
	}

	for _, fPath := range filePaths {
		fp, err := os.Open(fPath)
		if err != nil {
			log.Fatal(err)
		}
		countLines(fp, counts)
		fp.Close()
	}
	printCounts(out, counts)
}

func countLines(fp *os.File, counts map[string]fileNameMap) {
	s := bufio.NewScanner(fp)
	for s.Scan() {
		t := s.Text()
		if counts[t] == nil {
			counts[t] = make(fileNameMap)
		}
		counts[t][fp.Name()]++
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func printCounts(out io.Writer, counts map[string]fileNameMap) {
	for word, m := range counts {
		for fname, n := range m {
			if n > 1 {
				fmt.Fprintf(out, "%s: %d\t%s\n", word, n, fname)
			}
		}
	}
}
