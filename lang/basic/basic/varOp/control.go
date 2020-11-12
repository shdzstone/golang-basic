package varOp

import (
	"fmt"
)

func Control() {
	fmt.Println("控制语句用法-------------------")
	a := 4
	if a > 0 {
		fmt.Println("a大于0")
		if a > 3 {
			fmt.Println("a大于3")
		}
	} else {

		fmt.Println("a小于等于0")
	}

	var b interface{}
	b = 32
	switch b.(type) {
	case int:
		fmt.Println("a类型为整型")
	case string:
		fmt.Printf("a类型为字符串")
	default:
		fmt.Printf("以上条件均不满足!")
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("count:%d\n", i)
		if i < 3 {
			continue
		} else {
			break
		}
	}

	c := []string{"香蕉", "苹果", "雪梨"}

	for _, value := range c {
		//fmt.Println("key的值为",key)
		fmt.Println("value的值为", value)
	}

}
