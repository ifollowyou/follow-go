package kttp

import (
	"strings"
	"fmt"
)

const HTTP_PREFIX string = "http"
const HTTP_PROTOCOL string = "http://"

func CheckUrl(url string) (string) {
	// 判断 url 以加上前缀，可能是 https， 所以只判断http
	if !strings.HasPrefix(url, HTTP_PREFIX) {
		url = HTTP_PROTOCOL + url
	}
	fmt.Println(url)
	return url
}