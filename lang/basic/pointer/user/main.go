package main

import (
	"fmt"
)

/*
1.方法的值接收者和指针接收者receiver
* 方法传入接收者参数的类型并不重要，重要的是方法定义时receiver的类型是T还是*T。
* 若方法接收者为*T，则方法调用时，实参为T或*T没有区别，编译器会自动将形参转换为指针类型。
* 若方法接收者为T，则方法调用时，实参传为T或*T也没有区别，编译器会自动的将形参转换为值类型。
*/

func test1() {
	u1 := User{Name: "alex"}
	fmt.Printf("u1: %p\n", &u1)
	u1.SetName("jimmy")
	fmt.Println("u1.Name:", u1.Name) // alex

	u2 := &User{Name: "alex"} // u2为指针
	fmt.Printf("u2: %p\n", u2)
	u2.SetName("xxxxx")
	fmt.Printf("u2.name:%s\n", u2.Name)
	u2.SetName2("jimmy")
	fmt.Println("u2.Name:", u2.Name) // jimmy
}

func test2() {
	u1 := &User{Name: "alex"} // u1为指针
	fmt.Printf("u1: %p\n", &u1)
	u1.SetName("jimmy")
	fmt.Println("u1.Name:", u1.Name) // alex

	u2 := User{Name: "alex"}
	fmt.Printf("u2: %p\n", &u2)
	u2.SetName2("jimmy")
	fmt.Println("u2.Name:", u2.Name) // jimmy
}

func test3() {
	//u1 := pointer.User{Name: "alex"}
	// cannot use u (type User) as type *User in argument to ShowName
	//ShowName2(u1)

	//cannot use u2 (type *User) as type User in argument to ShowName
	//u2 := &pointer.User{Name: "jimmy"}
	//ShowName(u2)
}

func ShowName(u User) {
	fmt.Println("ShowName:", u.Name)
}

func ShowName2(u *User) {
	fmt.Println("ShowName:", u.Name)
}

func main() {
	//接收者测试
	test1()
	//test2()
	//test3()
}
