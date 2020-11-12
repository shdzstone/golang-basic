package main

import (
	"log"
)

/*
4.chan通信
* 通过通信来共享内存的确切涵义是：不同的协程间通过通道来通信
* 使用非缓冲通道阻塞任务结束
* 可以使用两个通道来处理协程通信
*/

//done：通过通信来通知外面事情做完了
//同一个goroutine同时绑定两个通道：in chan和done chan
func doWorker(id int, ch chan int, done chan bool) {
	for n := range ch {
		log.Printf("Worker goroutine %d received:%c \n", id, n)
		//要等两次，所以这个通道放在另一个goroutine中处理
		//go func() { done <- true }()
		done <- true
	}
}

//chan结构体，接收chan和完成chan，以便发送方与接收方通信
type worker struct {
	in   chan int
	done chan bool
}

func createdWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createdWorker(i)
	}

	//非缓冲通道发送都是阻塞式的：发一个数据过去都必须要被接收，不然会报deadlock
	for i, worker := range workers {
		log.Println("main goroutine send: ", i)
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		log.Println("main goroutine send: ", i)
		worker.in <- 'A' + i
	}

	//wait for all of them
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
