package varOp

import "fmt"

func Relational() {
	fmt.Println("关系运算符使用-----------------------------")

	a := 2
	b := 2

	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
	fmt.Printf("a==b:%t\n", a == b)
	fmt.Printf("a>b:%t\n", a > b)
	fmt.Printf("a<b:%t\n", a < b)
	fmt.Printf("a>=b:%t\n", a >= b)
	fmt.Printf("a<=b:%t\n", a <= b)
}
