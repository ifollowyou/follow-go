package main

import (
	"fmt"
	"strconv"
)

// 将数字转换为字母，使用 Sprintf 方法
func useSprint(val uint8) string {
	return fmt.Sprintf("%q", val)
}

// 将数字转换为字母，使用 Itoa 方法，暂未实现
func useItoA(val uint8) string {
	return strconv.Itoa(int(val))
}

func main() {
	var val uint8 = 65

	fmt.Println(useSprint(val))
	fmt.Println(useItoA(val))
}
