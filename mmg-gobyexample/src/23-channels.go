package main

import (
	"fmt"
	"time"
)

func main() {

	// 通道的创建语法 `make(chan val-type)`
	messages := make(chan string)

	fmt.Println(time.Now())

	// 向通道发送数据的语法为： `channel <-`
	// 下面的示例在一个匿名函数中发送数据。
	go func() {
		time.Sleep(time.Second * 5)
		messages <- "ping"
	}()

	// 从通道中读取数据的语法为： `<-channel`
	// 从通道中读取数据是同步阻塞的方式。
	msg := <-messages
	fmt.Println(msg)
	fmt.Println(time.Now())
}
