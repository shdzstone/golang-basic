package main

import (
	"fmt"
)

/*
1.panic
* 停止继续执行当前函数panic之后的代码
* 开始执行延时调用栈中的延时调用，从上至下，依次执行每一层的defer
* 若延时调用栈中没有defer延时调用，程序退出

2.recover
* 仅在defer中调用
* 用以捕获panic的值，并进行处理
* 若无法处理，可以重新panic，即repanic
*/

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			//panic(err)
			panic(fmt.Sprintf("I don't know what to do :%v", r))
		}
	}()
	//panic(errors.New("this is an error"))

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123)
}

func main() {
	tryRecover()
}
