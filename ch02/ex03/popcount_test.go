package popcount

import (
	"testing"
)

func TestZero(t *testing.T) {
	testZero(t, PopCount)
	testZero(t, PopCountWithLoop)
}

func testZero(t *testing.T, popCount func(uint64) int) {
	output := popCount(0)

	if output != 0 {
		t.Errorf("PopCount is %d, want 0", output)
	}
}

func TestAllBits(t *testing.T) {
	testAllBits(t, PopCount)
	testAllBits(t, PopCountWithLoop)
}

func testAllBits(t *testing.T, popCount func(uint64) int) {
	output := popCount(0xffffffffffffffff)

	if output != 64 {
		t.Errorf("PopCount is %d, want 64", output)
	}
}

func TestEachByte(t *testing.T) {
	testEachByte(t, PopCount)
	testEachByte(t, PopCountWithLoop)
}

func testEachByte(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 8; i++ {
		var value uint64 = 0xff << (uint(i) * 8)
		output := popCount(value)

		if output != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, output)
		}
	}
}

func Test0x5555(t *testing.T) {
	test0x5555(t, PopCount)
	test0x5555(t, PopCountWithLoop)
}

func test0x5555(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 4; i++ {
		var value uint64 = 0x5555 << (uint(i) * 8)
		output := popCount(value)

		if output != 8 {
			t.Errorf("PopCount(%x) is %d, want 8", value, output)
		}
	}
}

func TestEachOneBit(t *testing.T) {
	testEachOneBit(t, PopCount)
	testEachOneBit(t, PopCountWithLoop)
}

func testEachOneBit(t *testing.T, popCount func(uint64) int) {
	for i := 0; i < 64; i++ {
		var value uint64 = 1 << uint(i)
		output := popCount(value)

		if output != 1 {
			t.Errorf("PopCount(%x) is %d, want 1", value, output)
		}
	}
}

var output int

func BenchmarkPopCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(0x1234567890ABCDEF)
	}
	output = s
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithLoop(0x1234567890ABCDEF)
	}
	output = s
}

/* benchmark results.
➤ go test -bench  . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/nesheep5/study-gopl/ch02/ex03
BenchmarkPopCount-8           	2000000000	         1.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountWithLoop-8   	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/nesheep5/study-gopl/ch02/ex03	4.453s
*/
