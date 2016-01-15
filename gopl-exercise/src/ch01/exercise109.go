// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.

package main

import (
	"fmt"
	"net/http"
	"os"
	"ks/kttp"
)

const HTTP_PREFIX  string = "http://"



func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(kttp.CheckUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		println("resp.Status=", resp.Status)
		resp.Body.Close()
	}
}
