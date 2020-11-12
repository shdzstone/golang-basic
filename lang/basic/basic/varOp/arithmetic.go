package varOp

import (
	"fmt"
)

func Arithmetic() {
	fmt.Println("算术运算符-----------------------------")
	//变量赋值运算
	//:=只能用于函数体中的局部变量赋值
	a := 2
	b := 3
	fmt.Printf("a+b:%d\n", a+b)
	fmt.Printf("a-b:%d\n", a-b)
	//++只作为运算符，而不是表达式
	a++
	fmt.Printf("a++:%d\n", a)
	//--只作为运算符，而不是表达式
	b--
	fmt.Printf("b--:%d\n", b)
	fmt.Printf("a%%b:%d\n", a%b)
}
