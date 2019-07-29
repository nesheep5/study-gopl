package main

import (
	"io/ioutil"
	"testing"
)

func args(n int) []string {
	args := make([]string, 0, n)
	for i := 0; i < n; i++ {
		args = append(args, "arg")
	}
	return args

}

// go test -bench=. -benchmem

func BenchmarkEchoA(b *testing.B) {
	args := args(100)
	for i := 0; i < b.N; i++ {
		EchoA(ioutil.Discard, args)
	}
}

func BenchmarkEchoB(b *testing.B) {
	args := args(100)
	for i := 0; i < b.N; i++ {
		EchoB(ioutil.Discard, args)
	}
}
