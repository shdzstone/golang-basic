package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
1.函数的定义
* func eval(a,b int,op string)int
* 函数可以返回多个值，仅用于比较简单的函数
* 多返回时可以起名字，编译器自动生成(command+alt+v)接收函数返回值的变量
* 多返回值可以起名字，对于调用者而言没有区别，仅起提示作用
* 多返回值的常用场景是：第一个正常值、第二个error
* 匿名函数可作为函数参数使用，在调用时直接定义一个匿名函数来调用即可，GO没有lambda表达式
* GO语言没有默认参数、可选参数、函数重载、操作符重载
* GO语言只有可变参数列表
*/

//多返回值用法，一般是普通返回值+error
//函数error返回值方法
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation:" + op)
	}
}

//函数可以返回多个值
//13/3=4 ... 1
//函数返回多个值时可以起名字，使用编译器自动生成(command+alt+v)接收函数返回值的变量
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	//return q, r
	return a / b, a % b
}

//函数嵌套
func apply(op func(int, int) int, a, b int) int {
	//利用反射机制reflect,获取指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

//pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	//函数返回值error捕获
	if i, err := eval(3, 4, "*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	//函数值接收方式：定义多个变量，一对一接收函数返回值
	q, r := div(13, 3)
	fmt.Println(q, r)
	//可以使用_占位符，略去不接收的返回值
	a, _ := div(16, 5)
	fmt.Println(a)

	//函数嵌套
	//fmt.Println(apply(pow, 2, 4))
	//go语言没有lambda表达式，定义匿名函数即可
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	//可变参数列表
	fmt.Println(sum(1, 2, 3, 4, 5))
}
