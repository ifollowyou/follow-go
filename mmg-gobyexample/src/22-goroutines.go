package main

import "fmt"

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 直接调用方式
	f("direct")

	// goroutine调用方式
	go f("goroutine")

	// goroutine匿名调用方式
	go func(msg string) {
		f(msg)
	}("going")

	// 上面两个函数以goroutine方式异步执行
	// 当代码执行到Scanln时会暂停下来等待控制台的输入
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
