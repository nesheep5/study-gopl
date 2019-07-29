package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	out := new(bytes.Buffer)
	args := []string{"./echo", "arg1", "arg2", "arg3"}

	Echo(out, args)

	got := out.String()
	want := "arg1 arg2 arg3\n"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
