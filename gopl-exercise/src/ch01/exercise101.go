// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println
	// 全部参数
	p("Args: ", os.Args)
	// 程序名
	p("Args[0]: ", os.Args[0])
}