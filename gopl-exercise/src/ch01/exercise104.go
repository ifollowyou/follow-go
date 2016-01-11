// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io/ioutil"
	"log"
)

var pf = fmt.Printf

func main() {
	dup2()
}

// 原源码中的第一个实现，输出日志，循环部分作了退出修改
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	log.Print("Please input lines, and 'q' to done!")
	for input.Scan() {
		// 本循环是死循环，没有退出循环条件，一直等待输入
		line := input.Text();
		if line == "q" {
			break;
		}

		counts[line]++
	}
	// NOTE: ignoring potential errors from input.Err()

	log.Println("Counting...")
	for line, n := range counts {
		if n > 1 {
			pf("%d\t%s\n", n, line)
		}
	}

	log.Println("Finished...")
}

// 源码中的第二个实现，本次练习根据它进行修改
func dup2() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		// 循环处理文件
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}
}

// 处理单个文件内容
func countLines(f *os.File) (map[string]int) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, f.Name())
		}
	}

	return counts
}

// 源码中的第三个实现
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}