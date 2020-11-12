package main

import "fmt"

type User struct {
	Name string
}

func (x User) SetName(name string) {
	fmt.Printf("x1: %p\n", &x)
	x.Name = name
}
func (x *User) SetName2(name string) {
	fmt.Printf("x2: %p\n", x)
	x.Name = name
}
func (x User) GetName() string {
	return x.Name
}
