package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, "out."+strconv.Itoa(i))
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // ch チャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait()
}

func fetch(url string, ch chan<- string, ofile string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	wg.Add(1)
	go func() {
		if err := mkDir("out", 0777); err != nil {
			ch <- fmt.Sprint(err)
			return
		}

		file, err := os.Create("out/" + ofile)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		file.Write(b)
		file.Close()
		wg.Done()
	}()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, len(b), url)
}
func mkDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, mode)
	}
	return nil
}
