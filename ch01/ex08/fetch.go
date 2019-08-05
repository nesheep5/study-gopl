package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	// HTTPPrefix is Http prefix
	HTTPPrefix = "http://"
)

func main() {
	fetch(os.Stdout, os.Stderr, os.Args)
}

func fetch(out io.Writer, errout io.Writer, args []string) {

	url := args[1]
	if !strings.HasPrefix(url, HTTPPrefix) {
		url = HTTPPrefix + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(errout, "fetch: %v\n", err)
		os.Exit(1)
	}
	_, err = io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(errout, "fetch: reading %s: %v\n", url, err)
	}
}
