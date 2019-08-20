package popcount

import (
	"testing"
)

var output int

func BenchmarkPopCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(0x1234567890ABCDEF)
	}
	output = s
}

func BenchmarkPopCountByShifting(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountByShifting(0x1234567890ABCDEF)
	}
	output = s
}

/*
➤ go test -bench  . -benchmem
	goos: darwin
	goarch: amd64
	pkg: github.com/nesheep5/study-gopl/ch02/ex04
	BenchmarkPopCount-8             	2000000000	         1.39 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPopCountByShifting-8   	30000000	        47.8 ns/op	       0 B/op	       0 allocs/op
	PASS
	ok  	github.com/nesheep5/study-gopl/ch02/ex04	4.720s
*/
