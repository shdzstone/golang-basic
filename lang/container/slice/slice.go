package main

import "fmt"

/*
1.切片
type slice struct {
	array unsafe.Pointer // 元素指针，指向底层数组中slice初始元素位置
	len   int // 长度，表示切片中可用元素的个数，使用下标对slice的元素进行访问时，下标不能超过slice的长度；
	cap   int // 容量，底层数组的元素个数，容量 >= 长度。在底层数组不进行扩容的情况下，容量也是slice可以扩张的最大限度。
}
* go语言中一般不直接使用数组，通常使用切片slice
* 切片中的范围区间为左开右闭：比如arr[2:6]包含第2个元素，却不包含第6个元素，和数组中的count类似
* slice本身没有数据，是对底层array的一个view：数组的值改变，相应的slice值也会改变，slice改变，相应的数组值也会改变

2. Reslice
* slice支持reslice，即切片嵌套
* s := arr[2:6]
* s = s[:3]
* s = s[1:]

3.slice扩展
* 如下定义：
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 = arr[2:6]
s2 = s1[3:5]
* slice是可以向后扩展的，不可以向前扩展
* s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

4.slice添加元素
* slice添加元素时如果超越cap,系统会重新分配更大的底层数组，并且把原来的元素拷过去
* 由于值传递的关系，必须接收append的返回值,append时slice的len和cap都可能变掉
* s = append(s,val)

5.slice变量
* slice也是值类型，但slice结构体包含起始元素的指针、len、cap，所以slice传参后改变slice的值，原变量的值也会改变
*/

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("arr =", arr)
	fmt.Printf("s1 =%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 =%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("Adding elements to slice")
	s3 := append(s2, 10)
	//s4 and s5 no longer view arr,view different array
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s1,s2,s3 =", s3, s4, s5)
	fmt.Println("arr = ", arr)
}
