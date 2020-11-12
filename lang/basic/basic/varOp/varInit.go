package varOp

import (
	"fmt"
	"reflect"
)

//:=赋值格式不能在函数体外使用
func VarInit() {
	fmt.Println("变量赋值:=用法-----------------------------")

	m, _, n := 11, 12, 123
	fmt.Println(m)
	fmt.Println(n)

	//var i,j,k := 11,12,123
	i, j, k := 11, 12, 123
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(j)
	fmt.Println(k)
}
