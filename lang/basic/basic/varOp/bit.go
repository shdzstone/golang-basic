package varOp

import "fmt"

func Bit() {
	fmt.Println("位运算符------------------------")
	a := byte(0)
	b := byte(1)
	fmt.Printf("a=%b\n", a)
	fmt.Printf("b=%b\n", b)
	fmt.Printf("a&b:%d\n", a&b)
	fmt.Printf("b&b:%d\n", b&b)
	fmt.Printf("a|b:%d\n", a|b)
	fmt.Printf("a^b:%d\n", a^b)
	fmt.Printf("a^a:%d\n", a^a)
	fmt.Printf("b^b:%d\n", b^b)
	fmt.Printf("b<<1:%d\n", b<<1)
	fmt.Printf("b>>1:%d\n", b>>1)

}
