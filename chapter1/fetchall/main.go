package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, _ := os.Create("output.txt")
	for _, url := range os.Args[1:] {
		go fetch(url, ch, file) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // recieve from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, file *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//	defer resp.Body.Close() // Don't leak resources
	io.Copy(file, resp.Body)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
