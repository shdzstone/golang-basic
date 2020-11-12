package varOp

import (
	"fmt"
	"reflect"
)

const (
	cat string = "猫"
	dog        = "狗"
)
const apple, banana string = "苹果", "香蕉"
const a, b = 1, "香蕉"
const alen = len(b)

//常量定义使用
func Constant() {
	fmt.Println("常量定义及使用-----------------------------")
	fmt.Println(cat)
	fmt.Println(dog)
	fmt.Println(apple)
	fmt.Println(banana)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)
	fmt.Println(alen)
}
