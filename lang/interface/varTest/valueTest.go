package main

import "fmt"

/*
go语言中给接口赋值的时候
* 若接收者是是值，则指针和值均可以赋值给接口变量
* 若接收者是指针，则只能赋指针给接口变量，则可以自动实现值的处理，
*/

//定义Integer类型
type LessAddInf interface {
	Less(n Integer) bool
	Add(n Integer) Integer
}

type Integer int

func (i Integer) Less(n Integer) bool {
	return i < n
}

func (i *Integer) Add(n Integer) Integer {
	*i += n
	return *i
}

type Thing interface {
	Name() string
	Attribute() string
}

type Computer struct {
	CPU    string "计算器"
	Memory string "内存"
}

func (c Computer) Name() string {
	return "Computer"
}

func (c *Computer) Attribute() string {
	return fmt.Sprintf("CPU=%v Memory=%v", c.CPU, c.Memory)
}

func main() {
	var inf LessAddInf //interface变量
	var n Integer      //integer变量
	inf = &n
	fmt.Printf("inf(%v).Less(20)=%v\n", inf, inf.Less(20))
	fmt.Printf("inf(%v).Add(30)=%v\n", inf, inf.Add(30))

	var thing Thing
	var computer = Computer{CPU: "英特尔至强-v3440", Memory: "三星DDR4(8g)"}
	thing = &computer
	fmt.Printf("thing.Name()=%v\n", thing.Name())
	fmt.Printf("thing.Attribute()=%v\n", thing.Attribute())
}
