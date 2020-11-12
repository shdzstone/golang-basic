package varOp

import (
	"fmt"
	"unsafe"
)

//数据类型测试
func VarType() {
	fmt.Println("数据类型使用-----------------------------")

	var i int = 1
	fmt.Println(unsafe.Sizeof(i))
	var j uint32 = 1
	fmt.Println(unsafe.Sizeof(j))
	var k bool = true
	fmt.Println(k)
	fmt.Println(unsafe.Sizeof(k))
}
