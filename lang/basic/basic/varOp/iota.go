package varOp

import "fmt"

var Name = "hello ping"

/*
iota规则：
* 只能在常量表达式中使用
* 每次const出现时，都会让iota初始化为0
* 仅定义常量名而不定义常量值时，常量值自动为iota+1
* 可以使用下划线跳过不想要的值,但iota也会自增
* 可以使用运算符
* 可以跳过赋值的常量
*/

const (
	iotaA = iota
	iotaB = 3.14
	_
	iotaC = iota * 2
	iotaD
	iotaE
	iotaI, iotaJ = iota, iota + 1
	iotaM, iotaN
)

func Iota() {
	fmt.Println("iota使用-----------------------------")
	fmt.Print("iotaA：")
	fmt.Println(iotaA)
	fmt.Print("iotaB：")
	fmt.Println(iotaB)
	fmt.Print("iotaC：")
	fmt.Println(iotaC)
	fmt.Print("iotaD：")
	fmt.Println(iotaD)
	fmt.Print("iotaE：")
	fmt.Println(iotaE)
	fmt.Print("iotaI：")
	fmt.Println(iotaI)
	fmt.Print("iotaJ：")
	fmt.Println(iotaJ)
	fmt.Print("iotaM：")
	fmt.Println(iotaM)
	fmt.Print("iotaN：")
	fmt.Println(iotaN)
}
