package main

import "fmt"

/**
1.slice定义
* 通常slice不由array初始化，可以直接定义
* slice扩充时，cap扩容通常乘以2

2.slice删除元素
* 通过数组中元素索引，删除
* 首元素和尾元素删除

3.slice拷贝
* 通过copy函数实现
*/

func printSlice(s []int) {
	fmt.Printf("value=%v len=%d cap=%d\n", s, len(s), cap(s))
}
func main() {
	fmt.Println("Creating slice")
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16) //创建指定长度的slice
	printSlice(s2)
	s3 := make([]int, 16, 32) //创建指定长度len和容量cap的slice
	printSlice(s3)

	//slice拷贝
	fmt.Println("Copying slice")
	copy(s2, s1) //slice拷贝
	printSlice(s2)

	//删除slice中间元素
	fmt.Println("Deleting elements from slice")
	//slice没有直接删除的语法,可以通过append来实现
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	//	slice = append([]byte("hello "), "world"...)
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	//删除slice头元素
	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	//删除slice尾元素
	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	fmt.Println(tail)
	s2 = s2[:len(s2)-1]
	printSlice(s2)
}
