package main

import (
	"fmt"
	"golang-basic/lang/queue"
)

func main() {
	q := queue.Queue{1}
	//可以直接改变q的值
	q.Push("a")
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsNil())
	fmt.Println(q.Pop())
	fmt.Println(q.IsNil())
}
