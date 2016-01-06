// 使用 select+default 来实现无阻塞操作。
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println(time.Now())

	// 如果没有缓冲，则全部 select 都走 default 分支。
	messages := make(chan string, 1)
	signals := make(chan bool)

	// 在下面的示例中，如果 messages 中有数据，则输出该数据；
	// 否则直接走 default 分支退出。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	log.Println(time.Now())

	// A non-blocking send works similarly.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	log.Println(time.Now())

	// 多 case 的使用方式一致。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	log.Println(time.Now())
}
