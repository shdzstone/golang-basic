package main

import (
	"fmt"
	"time"
)

/*
1.缓冲通道
* 只收不发panic
* 多发少收正常运行
* 少发多收panic
* 同一goroutine内部，代码的执行遵循从上到下的顺序
* 缓冲通道的发送操作与接收操作处于同一goroutine时，相当于管子的两头都在同一个goroutine，
* 因为有缓冲，发送操作与接收操作可以不必同步执行
* 缓冲通道的发送操作与接收操作可以处于同一goroutine，即可以非并发的操作
*/

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	time.Sleep(time.Millisecond)
}
