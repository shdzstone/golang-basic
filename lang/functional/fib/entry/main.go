package main

import (
	"bufio"
	"fmt"
	"golang-basic/lang/functional/fib"
	"io"
)

/*
1.闭包使用实例
* 斐波那契数列
* 为函数实现接口
* 使用函数来遍历二叉树
* 可以使用函数类型来声明匿名函数的类型
* 函数类型也可以实现接口
*/

func printFileContents(reader io.Reader) {
	//bufio.NewScanner一行一行的扫描内容
	//file实现了Reader接口，因此参数是io.Reader接口，而不是file对象
	scanner := bufio.NewScanner(reader)
	//省略初始条件和递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	printFileContents(f)
}
