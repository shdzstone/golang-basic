package main

import "fmt"

/**
1.数组定义
* 数量写在类型之前
* var arr [5]int					//空数组
* arr2 := [3]int{1,2,5}				//定长数组
* arr3 := [...]int { 2,4,6,8,10}	//可变长数组
* var arr4 [4][5]int				//二维数组

2.数组遍历
* len()及range的使用
* for i := 0; i < len(arr3); i++ {}
* for _, i := range arr3 {}
* 可通过"_"占位符省略变量
* 不仅是range，任何地方都可以使用"_"省略变量
* 如果只要i，可写成for i:=range arr {}

3.为什么使用range
* 意义明确，美观
* c++:没有类似能力
* java/python:只能通过for each value，不能同时获得i,v

4.数组是值类型
* [3]int和[5]int是不同的类型
* go语言中，数组是值类型，函数传数组参数时，函数体对数组的修改不会影响原数组
* 调用func f(arr [10]int)会"拷贝"数组
* 大部分编程语言，数组均为引用类型，函数传参为引用传递
* 在go语言中一般不直接使用数组，真正使用的是切片slice
*/

//数组定义、初始化及遍历
func array_init() {

	var arr [5]int
	arr2 := [3]int{1, 2, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var arr4 [4][5]int

	fmt.Println(arr, arr2, arr3)
	fmt.Println(arr4)

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	//range关键字可以同时获得数组元素的位置及value
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, i := range arr3 {
		fmt.Println(i)
	}
}

func maxVal(arr [5]int) (maxi int, maxValue int) {
	maxi = -1
	maxValue = -1

	for i, v := range arr {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	return maxi, maxValue
}

//传递指针，不拷贝
func printArray(arr []int) {
	//值传递，拷贝数组
	//func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr3 := [...]int{2, 4, 6, 8, 10}
	maxi, maxValue := maxVal(arr3)
	fmt.Println(maxi, maxValue)

	printArray(arr3[:])
	fmt.Println(arr3)
}
