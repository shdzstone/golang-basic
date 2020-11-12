package _var

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内部的全局变量
var (
	AA = 3
	SS = "kkk"
	BB = true
)

/*
1. 使用var关键字声明nil变量或有初值的变量
* var a,b,c bool
* var s1,s2 string = "hello","world"
* 可放在函数内，或直接放在包内
* 使用var()集中定义变量

2. 让编译器自动决定类型
* var a,b,i,s1,s2 = true,false,1,"hello","world"

3. 使用:=定义变量
* a,b,i,s1,s2 = true,false,1,"hello","world"
* 只能在函数内使用，不能定义全局变量，也不能定义包变量

4. 内建变量类型
* bool,string
* (u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
* byte,rune(字符类型)
* float32,float64,complex64,complex128

5. 复数回顾
* z=a+bi
* 欧拉公式

6. 强制类型转换
* 类型转换是强制的
* var a,b int =3,4
* var c = math.Sqrt(float64(a*a + b*b))  xxx
* var c = int(math.Sqrt(float64(a*a + b*b))) √
* 浮点数不准确的问题

7. 常量
* const filename = "abc.txt"
* 常量可以定义在函数体
* 常量可以定义为包变量
* const()可以定义多个常量
* const数值可以作为各种类型使用
* const a,b = 3,4
* var c int = int(math.Sqrt(a*a+b*b))

8.枚举
* 普通枚举类型
* 自增值枚举类型，iota的使用
* _占位符的使用


9.总结：变量定义的要点
* 变量类型写在变量名之后
* 编译器可推测变量类型
* 没有char，只有rune(32位)
* 原生支持复数类型
*/

//:=不能用在函数体外
// var a := 1

//变量类型默认初始值
func VariableZeroValue() {
	var a int
	var s string
	//%q可以打印出空串
	fmt.Printf("%d %q\n", a, s)
}

//变量声明
func VariableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//类型预测
func VariableTypeDeduction() {
	var a, b, c, d = 3, 4, "string", true
	s := "abc"
	a = 5
	fmt.Println(a, b, c, d, s)
}

//声明、赋值、初始化简写
func VariableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//使用复数验证欧拉公式
func Euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	//浮点数有位数限制，不准确
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//强制类型转换
func Triangle() {
	var a, b int = 3, 4
	fmt.Println(CalcTriangle(a, b))
}

func CalcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量
func Consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

//枚举
func Enums() {
	const (
		cpp = iota
		_
		python
		golang
		java
	)
	fmt.Println(cpp, python, golang, java)
	//b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
