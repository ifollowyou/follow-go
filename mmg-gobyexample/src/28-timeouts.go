// 超时机制对需要远程调用的程序来说是很重要事情，
// 在 Go 语言中，使用 chan 和 select 可以很方便地实现超时处理。
package main

import "time"
import "fmt"
import "log"

func main() {
	log.Println(time.Now())
	// 假定需要远程调用处理需要2s才有响应
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 由 select+time.After 实现的超时处理机制，只要有 case 响应，select 就会停止阻塞。
	// 下面的例子超时时间为1秒，到时以后不再等待c1的响应数据。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	log.Println(time.Now())

	// 下面的示例超时时间为3秒，而任务处理只需2秒，正常输出任务响应。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	log.Println(time.Now())
}
