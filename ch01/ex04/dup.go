package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	Dup2(os.Stdin, os.Stdout, os.Args)

}

// Dup2 count up and display filename to duplicate words in files.(Stream Mode)
func Dup2(in *os.File, out io.Writer, args []string) {
	counts := make(map[string]int)
	fileNames := make(map[string]string)
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
		countLines(fp, counts, fileNames)
		fp.Close()
	}
	printCounts(out, counts, fileNames)
}

func countLines(fp *os.File, counts map[string]int, fileNames map[string]string) {
	s := bufio.NewScanner(fp)
	for s.Scan() {
		t := s.Text()
		counts[t]++
		if counts[t] == 1 {
			fileNames[t] = fp.Name()
		} else {
			if !strings.Contains(fileNames[t], fp.Name()) {
				fileNames[t] += " " + fp.Name()
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func printCounts(out io.Writer, counts map[string]int, fileNames map[string]string) {
	for word, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%s: %d\t%s\n", word, n, fileNames[word])
		}
	}
}
