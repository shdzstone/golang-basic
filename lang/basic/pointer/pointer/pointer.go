package pointer

import "fmt"

/**
1.指针
* var a *int代表int型指针a，涵义为a是一个内存地址，可以用取值操作符*,来取指针a中的值
* *为取值运算符,根据指针访问变量的值(内存实体)
* &为取址运算符,根据变量的值(内存实体)获取变量的内存地址
* &a引用传递，返回相应变量a的内存地址
* go中的指针不能运算，不能指针加减，即p *array= array[i],p++之类的操作
* 值传递：给函数传参时，为参数做一份拷贝，将参数的值赋值给拷贝，函数体中修改拷贝的值，原参数的值不会改变
* 引用传递：给函数传参时，不给原参数做拷贝，直接将参数的内存地址传给函数，函数中可直接修改原参数的值
* java、python中，一般自定义类型均为引用传递，系统内建类型（int、bool、float）为值传递
* go语言，只有值传递一种方式，凡是调函数，参数都要去拷一份
*/

func ChangeValue() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

//传递指针
func Swap(a, b *int) {
	*b, *a = *a, *b
}

//利用函数返回值赋值的特性，来实现原对参数的值的改变
func SwapNext(a, b int) (int, int) {
	return b, a
}

func main() {
	//指针测试
	ChangeValue()

	//交换两个值
	a, b := 3, 4
	//a, b = pointer.SwapNext(a, b)
	Swap(&a, &b)
	fmt.Println(a, b)
}
