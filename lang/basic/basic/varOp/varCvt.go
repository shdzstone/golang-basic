package varOp

import (
	"fmt"
	"reflect"
)

//类型转换
//i,j,k := 11,12,123
//var声明格式可以
//var i,j,k := 11,12,123
func VarCvt() {
	fmt.Println("类型转换-----------------------------")
	var a int
	var b int = 456
	var (
		c string = "stone:"
		d int    = 1987
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e := float32(b)
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))
}
