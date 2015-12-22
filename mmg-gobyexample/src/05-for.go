// `for` 是Go语言中唯一的循环结构，它有三种基本格式。

package main

import "fmt"

func main() {
	// 最基本的单条件循环格式
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的initial/condition/after格式
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 无条件死循环格式，通过break或return退出循环。
	for {
		fmt.Println("loop")
		break
	}
}
