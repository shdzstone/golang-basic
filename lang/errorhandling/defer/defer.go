package main

import (
	"bufio"
	"fmt"
	"golang-basic/lang/functional/fib"
	"os"
)

/*
1.defer
* defer常被用于关闭文件描述符、关闭数据库连接以及解锁资源
* 确保defer语句在当前函数或方法结束之前执行defer语句
* 确切的讲：defer语句执行时间是在return之后，但是在返回值返回给调用方之前执行，所以使用defer可以达到修改返回值的目的。
* defer列表的调用顺序，类似于栈的后进先出
* 参数在defer语句时计算

2.资源管理
* 资源的管理一般都是打开、关闭成对出现的
* 资源打开时就要考虑关闭文件，此时可以使用defer来延时关闭文件
* 创建了bufio，就要考虑flush

3. 何时使用defer
* Open/Close
* Lock/Unlock
* PrintHeader/PrintFooter

4. 错误处理
* go程序出现错误后会直接崩溃，crash掉
* 要进行error捕获，打印出错误原因，并return掉
* error是一个接口类型，可以自定义自己的error
* 一般文档均会给出err错误类型，可以捕获掉已知类型的error，未知类型的error可以简单打印并return掉

5. 如何实现统一的错误处理逻辑
*/

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("Printed too many")
		}
	}
}

func writeFile(filename string) {
	//创建文件
	//file, err := os.Create(filename)
	//If there is an error, it will be of type *PathError.
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//自定义error
	//err = errors.New("This is a custom error")
	//错误处理
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	//关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	//tryDefer()

	//writeFile("functional/fib/fib.txt")
	fmt.Println(f())
}
