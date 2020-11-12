package main

import (
	"fmt"
	"time"
)

/*
1.goroutine同步
* 新开的goroutine与主goroutine，要实现同步，可以使用非缓冲chan的阻塞接收特性
*/

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	fmt.Println("啦啦啦")
}

func main() {

	done := make(chan bool)
	go worker(done)

	<-done
}
