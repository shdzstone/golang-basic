package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1.sync标准库中的atomic包
* 系统atomic原子化操作函数，负责多个goroutine间数据的原子化操作
* atomic系统函数：atomic.AddInt32()/atomic.AddInt64()/atomic.AddUInt32()

2.sync标准库中的mutex包
*使用sync.Mutex变量的lock()/unlock()来模拟atomic.AddInt32()等操作
*/

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	//scope中锁的使用
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
