// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const HTTP_PREFIX  string = "http://"

// go run exercise108.go gopl.io
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

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
