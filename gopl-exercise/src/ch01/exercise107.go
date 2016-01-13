// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
// buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
)

// 本练习的 io.Copy 方式
func ioCopy(resp *http.Response) {
	return io.Copy(os.Stdout, resp)
}

// 原来的 ioutil.ReadAll 方式
func readAll(resp *http.Response) {
	return ioutil.ReadAll(resp.Body)
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioCopy(resp)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// go run exercise107.go http://gopl.io