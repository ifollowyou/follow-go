package main

import "fmt"

// 求和函数。函数名后先是入参，再是出参。
func plus(a int, b int) int {

	// 在Go语言中需要显式地返回，
	// 它并不能像Groovy一样自动返回末尾的表达式。
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
