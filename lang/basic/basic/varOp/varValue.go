package varOp

import (
	"fmt"
	"reflect"
	"unsafe"
)

type stoneInt int32

var varI int32
var varJ float32
var varB bool
var varD complex64
var varS string

func init() {
	fmt.Println("varValue初始化init函数----------------------")
	varI = 123123
}

//数据默认值测试
func VarValue() {
	fmt.Println("基本数据类型默认值----------------------------")

	fmt.Print("int32:")
	fmt.Println(varI)
	fmt.Print("float32:")
	fmt.Println(varJ)
	fmt.Print("bool:")
	fmt.Println(varB)
	fmt.Print("complex64:")
	fmt.Println(varD)
	fmt.Print("string:")
	fmt.Println(varS)
	fmt.Print("s类型:")
	fmt.Println(reflect.TypeOf(varS))
	var sint stoneInt
	fmt.Print("sint类型:")
	fmt.Println(reflect.TypeOf(sint))
	fmt.Print("sint内存大小:")
	fmt.Println(unsafe.Sizeof(sint))
}
