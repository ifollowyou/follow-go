// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
)

var p = fmt.Println

func byCount(args []string)  {
	p("Using count...")
	for count := 0; count<len(args); count++ {
		p("index=", count, "; value=" + args[count])
	}
}

func byRange(args []string)  {
	p("Using range...")
	for i, v := range args {
		p("index=", i, "; value=" + v)
	}
}

func main() {
	byRange(os.Args)
	byCount(os.Args)
}