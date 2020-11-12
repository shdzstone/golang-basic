package main

import (
	"fmt"
	"sync"
)

/*
3.sync标准库中sync.WaitGroup
* 使用sync.WaitGroup来做线程阴塞/同步
* sync.WaitGroup.Add(10)在某goroutine中添加需要等待的协程数量
* sync.WaitGroup.Done()其它协程通知给控制协程，当前执行完毕
* sync.WaitGroup.Wait()控制协程阻塞并等待其它协程处理结束
*/

//操作协程，并循环接收in通道过来的值
func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received:%c \n", id, n)
		//每接收完一个数据执行一个函数
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createdWorker(id int, wg *sync.WaitGroup) worker {
	//创建结构体
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	//创建worker
	for i := 0; i < 10; i++ {
		workers[i] = createdWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
