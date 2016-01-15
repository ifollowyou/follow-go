// Exercise 1.10: Find a website that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to print its out put to a file so it can be
// examined.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"ks/kttp"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Need a argument!")
		return
	}

	// 对于 slice 数组，使用 append 时必须要赋值回去
	results := []string{}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 开启线程
	}
	for range os.Args[1:] {
		result := strings.TrimSpace(<-ch)

		if len(result) > 0 {
			results = append(results, result) // 先保存到数组
		}
	}

	total := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	results = append(results, total)

	fmt.Println(results)
	write2file(results)
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(kttp.CheckUrl(url))
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func write2file(results []string) {
	now := time.Now().Nanosecond();
	// 通过文件对象更好地控制文件的读写
	name := strconv.Itoa(now) + ".txt"
	fmt.Println("name=" + name)
	f, err := os.Create(name)
	checkX(err)
	// 应当习惯性地在得到文件对象后，开启它的完成后自动关闭
	defer f.Close()

	for _, line := range results {
		f.WriteString(line + "\n")
	}
}

func checkX(e error) {
	if e != nil {
		panic(e)
	}
}

// go run ch01/exercise110.go baidu.com gopl.io
