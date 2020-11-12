package varOp

import "fmt"

func Logic() {
	a := true
	b := false
	c := true

	fmt.Printf("a=%t\n", a)
	fmt.Printf("a=%t\n", b)
	fmt.Printf("a=%t\n", c)

	fmt.Printf("a&&b:%t\n", a && b)
	fmt.Printf("a&&c:%t\n", a && c)
	fmt.Printf("a||b:%t\n", a || b)
	fmt.Printf("!a:%t\n", !a)

}
