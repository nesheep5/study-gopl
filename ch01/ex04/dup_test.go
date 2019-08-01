package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDup2(t *testing.T) {
	// TODO refactor the testfile creation func
	// testfileA
	textA := `aaa
bbb
bbb
zzz
aaa`

	fileA, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(fileA.Name()) // clean up

	if _, err := fileA.Write([]byte(textA)); err != nil {
		log.Fatal(err)
	}
	if err := fileA.Close(); err != nil {
		log.Fatal(err)
	}

	// testfileB
	textB := `bbb
bbb
ccc
ccc`

	fileB, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(fileB.Name()) // clean up

	if _, err := fileB.Write([]byte(textB)); err != nil {
		log.Fatal(err)
	}
	if err := fileB.Close(); err != nil {
		log.Fatal(err)
	}

	out := &bytes.Buffer{}
	Dup2(os.Stdin, out, []string{".Dup2", fileA.Name(), fileB.Name()})

	gots := strings.Split(out.String(), "\n")
	if !contains(gots, "aaa: 2\t"+fileA.Name()) {
		t.Errorf("aaa should be output")
	}
	if !contains(gots, "bbb: 2\t"+fileA.Name()) {
		t.Errorf("bbb should be output")
	}
	if !contains(gots, "bbb: 2\t"+fileB.Name()) {
		t.Errorf("bbb should be output")
	}
	if !contains(gots, "ccc: 2\t"+fileB.Name()) {
		t.Errorf("ccc should be output")
	}
	if contains(gots, "zzz: 2\t"+fileA.Name()) {
		t.Errorf("zzz should not be output")
	}
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
