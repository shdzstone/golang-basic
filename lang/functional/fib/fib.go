package fib

import (
	"fmt"
	"io"
	"strings"
)

/*
1.函数类型的使用
* 函数类型也能实现接口
* 函数是一等公民，函数自然也可以做为接收者
* 接口实现所需的receiver只不过是一个参数而已，只不过写在函数名字的前面
*/

//函数类型
type intGen func() int

//函数intGen实现Read接口
func (g intGen) Read(p []byte) (n int, err error) {
	//获取下一个斐波那契数
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	//将next转换为字符串
	s := fmt.Sprintf("%d\n", next)

	//TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

//1,1,2,3,5,8,13,21,34
//a,b
//  a,b
//    a,b
//斐波那契数列生成器
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
