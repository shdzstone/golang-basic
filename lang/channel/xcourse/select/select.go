package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1.select...default
* 除default外，若只有一个case语句评估通过，那么就执行这个case里的语句；
* 除default外，若有多个case语句评估通过，那么以伪随机的方式随机选一个；
* 若default 外的case语句都没有通过评估，那么执行 default 里的语句；
* 若没有default，那么代码块会被阻塞，直到有一个 case 通过评估；否则一直阻塞；

2.定时器的使用
* time.Sleep()/time.After()/time.Duration()/time.Tick()

3.在select使用Nil channel：阻塞式获取chan的值

4.通过通信来共享内存
* demo演示了使用chan来共享不同goroutine间的数据通信
* 不同goroutine间数据的传递未使用lock/wait等同步机制，却并发的解决了不同goroutine间的通信，典型的CSP使用实例

5.传统的同步机制
* WaitGroup
* Mutex

*/

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, ch chan int) {
	for n := range ch {
		time.Sleep(time.Millisecond)
		fmt.Printf("Worker %d received:%d \n", id, n)
	}
}

func createdWorker(id int) chan<- int {
	ch := make(chan int)
	go worker(id, ch)
	return ch
}

func main() {
	//数据生产者chan
	var c1, c2 = generator(), generator()
	//数据接收者chan
	w := createdWorker(0)

	//数据存放队列
	var values []int

	/*
		//time.After超时控制，从当前时间算起，隔段时间后，返回一个只读通道
		tm := time.After(10 * time.Second)
		//time.Tick用以间隔某段时间执行某任务，返回一个只读通道
		tick := time.Tick(time.Second)
	*/
	for {
		//nil信道阻塞特性的使用
		var activeValue int
		var activeWorker chan<- int

		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = w
		}

		select {
		//两个信道收到的数据排队
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			//nil信道也可以select到，但会阻塞
			//结果是有值时才插入到队列
			//数据阻塞式分发
		case activeWorker <- activeValue:
			values = values[1:]
			/*
				//定时器：每隔一秒查询下队列的状态
				case <-tick:
					fmt.Println("queue len =  ", len(values))
				//队列状态：每隔800毫秒看下是否超时
				case <-time.After(800 * time.Millisecond):
					fmt.Println("timeout")

				//延时器：10秒钟后结束接收数据
				case now := <-tm:
					fmt.Println("time:", now)
					return
			*/
		}

	}
}
