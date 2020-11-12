package main

import (
	"fmt"
	"time"
)

/*
1.chan
* chan信道是goroutine与goroutine之间的通信管道。
* chan和函数一样，也是一等公民，也可以做为变量、参数、返回值
* chan相当于一根管子，一头连接着发送goroutine，一头连接着接收goroutine

2.chan
* 操作符<-用来指定管道的方向，发送或接收。如果未指定方向，则为双向管道。
* 无缓冲信道chan相当于一个没有长度的管子，接收到的数据必须马上被消耗掉，即必须在chan接收到数据的同时被某个接通的goroutine消耗掉
* 无缓冲信道从不存储数据，只负责数据的流通，流入的数据必须要马上流出才可以。
* 无缓冲信道不能承载数据。非缓冲信道上若发生了有流入无流出，或者有流出无流入，会导致死锁。
* 所有向非缓冲信道chan存取数据的goroutine必须要成对：一定要一个线里存数据，一个线里取数据
* 无缓冲信道是一批数据一个一个的「流进流出」
* 向chan发送数据时，发送方和接收者必须同时都在运行，否则会阻塞

3.close chan
* Channel支持close操作，用于关闭channel，后面对该channel的任何发送操作都将导致panic异常。
* 对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据。
* 必须由发送方来关闭信道：chan close
* 接收方可通过,ok=<-ch,ok返回值来判断信道chan是否关闭
* 接收方也可通过range遍历来判断信道是否关闭

* buffered channel
* 缓存信道(buffered channel)意思是，缓冲信道不仅可以流通数据，还可以缓存数据。
* 缓存信道是有容量的，存入一个数据的话 , 可以先放在信道里，不必阻塞当前线而等待该数据取走。
* 当缓冲信道达到满的状态的时候，就会表现出阻塞了，因为这时再也不能承载更多的数据了。
* 缓冲信道会在满容量的时候加锁。
* 其实，缓冲信道是先进先出的，可以把缓冲信道看作为一个线程安全的队列
* buffered chan相当于一个有长度的管子，本身可以存储某长度的接收数据，长度为n，则可以存储n个数据
* 缓冲信道则是一个一个存储，然后一起流出去

3.chan理念
* 理论基础：Communication Sequential Process(CSP)
* Don't communicate by sharing memory; share memory by communicating;
* 不要通过共享内存来通信；通过通信来共享内存
*/

func worker(id int, ch chan int) {
	//通过range来判断chan是否close
	for n := range ch {
		//通过<-ch值是否为空判断chan是否close
		//判断<-ch发送的值是否为空
		//n, ok := <-ch
		//if !ok {
		//	break
		//}

		//print涉及io，需要go调度器调度，因此goroutine不一定顺序执行
		fmt.Printf("Worker %d received:%d \n", id, n)

	}
}

//创建chan并绑定goroutine
func createdWorker(id int) chan<- int {
	ch := make(chan int)
	//创建goroutine
	go worker(id, ch)
	return ch
}

//创建chan并绑定相应的goroutine
func chanDemo() {
	//var声明的chan，默认为nil，使用时需初始化
	//var ch chan int

	//:=来声明并初始化chan

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)

		//开10个协程，并绑定相应的chan
		//go worker(i, channels[i])

		channels[i] = createdWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

//单个chan相当于一个没有长度的管子，接收到的数据必须马上被消耗掉，即必须在chan接收到数据的同时被某个接通的goroutine消耗掉
//buffered chan相当于一个有长度的管子，本身可以存储某长度的接收数据，长度为n，则可以存储n个数据
func bufferedChannel() {
	ch := make(chan int, 3)
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	time.Sleep(time.Millisecond)
}

//chan关闭掉后goroutine还可以接收到数据，收到的数据是chan类型的空值,比如chan int：0
//chan close一定是发送方来关闭
//接收方判断接收方是否关闭：1.通过,ok:=接收到的值是否为空 2.通过range遍历
func channelClose() {
	ch := make(chan int)
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	close(ch)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
