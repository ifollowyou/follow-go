// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const HTTP_PREFIX  string = "http://"

// go run exercise109.go gopl.io
func checkUrl(url string) (string) {
	// 判断 url 以加上前缀
	if !strings.HasPrefix(url, HTTP_PREFIX) {
		url = HTTP_PREFIX + url
	}
	fmt.Println(url)
	return url
}

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(checkUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		println("resp.Status=", resp.Status)
		resp.Body.Close()
	}
}
