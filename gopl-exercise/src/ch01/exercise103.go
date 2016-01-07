// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var p = fmt.Println

func byConcat(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}

	return s;
}

func byJoin(args []string) string {
	return strings.Join(os.Args[1:], " ")
}

// 本程序测试的结果，跑100w次，使用join的效率更高时间更少
func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		byJoin(os.Args)
	}
	period := time.Since(start).Nanoseconds()
	p(period)

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		byConcat(os.Args)
	}
	period = time.Since(start).Nanoseconds()
	p(period)
}