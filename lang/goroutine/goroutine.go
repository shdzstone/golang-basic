package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
1.协程
* coroutine是轻量级的"线程"，比线程更小的执行单位
* coroutine是底层非抢占式多任务处理，由协程主动交出控制权
* 编译器/解释器/虚拟机层面的多任务，而不是操作系统层面的多任务，操作系统没有协程
* 线程是抢占式多任务处理，由系统决定什么时候切换控制权

2.goroutine
* goroutine是Go中最基本的执行/调度单元
* 每一个Go程序至少有一个goroutine：主goroutine。程序启动时，它会自动创建。
* 任何函数只需要加上go关键字就能送给调度器运行,相当于新开一个goroutine。
* 不需要在定义时区分是否是异步函数
* Goroutine支持异步抢占式切换，调度器负责goroutine间的切换
* go语言本身有一个协程调度器，用来调度协程的切换
* 多个goroutine可以一个线程或多个线程上运行，由调度器决定
* 使用-race来检测数据访问冲突

3.goroutine可能的切换点
* I/O，select
* channel
* 等待锁
* 函数调用(有时)
* runtime.Gosched()
* 上述只是参考，不能保证切换，不能保证在其它地方不切换

3.goroutine data race解决
* 终端切换到当前协程所在文件夹
* go run -race goroutine.go #查看当前data race情况
*/

//main函数也是由调度器开辟的一个goroutine
func main() {
	var b [10]int
	for i := 0; i < 10000; i++ {
		//此处的i调用传值过来的i,若直接使用外层的i,会造成data race
		//此处是一个闭包，直接引用外部i，当i++变成10之后，会造成a越界
		go func(a int) {
			for {
				//print是一引系统级的io操作，中间涉及到协程切换
				fmt.Printf("Hello from "+"goroutine %d \n", a)

				//a[i]++
				//手动交出goroutine控制权，让别的协程也有机会运行，再由调度器来调度运行本协程
				runtime.Gosched()
			}
		}(i)
	}
	//go语言中main函数执行完毕所有协程就随之结束，要想看到上述协程的执行结果，需要本协程休眠下
	time.Sleep(time.Millisecond)

	fmt.Println(b)
}
