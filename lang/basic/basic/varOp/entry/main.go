//package
package main

import (
	"fmt"
	"golang-basic/lang/basic/basic/varOp"

	"runtime"
)

//初始化函数测试
func init() {
	fmt.Println("Main init调用其它包的全局变量----------------------")
	fmt.Println(varOp.Name)
}

//程序执行入口
func main() {
	//数据类型
	varOp.VarValue()
	varOp.VarType()
	varOp.VarInit()
	varOp.VarCvt()
	varOp.Constant()
	varOp.Iota()
	//操作符
	varOp.Arithmetic()
	varOp.Relational()
	varOp.Logic()
	varOp.Bit()
	varOp.Control()
	fmt.Println(runtime.GOARCH)
}
