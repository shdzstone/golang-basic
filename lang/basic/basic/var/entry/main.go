package main

import (
	"fmt"
	"golang-basic/lang/basic/basic/var"
)

func main() {
	fmt.Println("hello")
	_var.VariableZeroValue()
	_var.VariableInitialValue()
	_var.VariableTypeDeduction()
	fmt.Println(_var.AA, _var.SS, _var.BB)
	_var.VariableShorter()
	_var.Euler()
	_var.Triangle()
	_var.Consts()
	_var.Enums()
}
