// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

/*
1.nil chan操作
* 向nil chan发送数据panic
* nil chan接收数据panic
* close() nil chan panic

2.非缓冲chan
* 只接不发会panic
* 只发不接会阻塞发送操作所在goroutine，但不会panic
* 成对的发与收，才能通畅运行
* 同一goroutine内部，代码的执行遵循从上到下的顺序
* 非缓冲通道的发送操作与接收操作处于同一goroutine时，相当于管子的两头都在同一个goroutine，
* 因为没有缓冲，发送操作与接收操作不能同时执行，会造成两个阻塞：发送方等待接收，接收方等待发送，从而造成死锁
*/

func main() {

	/*
		var msgTest chan string
		msgTest <- "123123"
		fmt.Println(<-msgTest)
		close(msgTest)
	*/

	/*
		messages := make(chan string)

		go func() {
			for {
				messages <- "ping"
			}

		}()

		go func() {
			for {
				msg := <-messages
				fmt.Println(msg)
			}
		}()

		time.Sleep(time.Millisecond)
	*/
	//read_only := make(<-chan int)

	//close(read_only)

	/*
		var c chan struct{} // nil
		select {
		case <-c: // 阻塞操作
		case c <- struct{}{}: // 阻塞操作
		default:
			fmt.Println("Go here.")
		}
	*/

	/*
		c := make(chan string, 2)
		trySend := func(v string) {
			select {
			case c <- v:
			default: // 如果c的缓冲已满，则执行默认分支。
			}
		}
		tryReceive := func() string {
			select {
			case v := <-c:
				return v
			default:
				return "-" // 如果c的缓冲为空，则执行默认分支。
			}
		}
		trySend("Hello!") // 发送成功
		trySend("Hi!")    // 发送成功
		trySend("Bye!")   // 发送失败，但不会阻塞。
		// 下面这两行将接收成功。
		fmt.Println(tryReceive()) // Hello!
		fmt.Println(tryReceive()) // Hi!
		// 下面这行将接收失败。
		fmt.Println(tryReceive()) // -
	*/
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
}
