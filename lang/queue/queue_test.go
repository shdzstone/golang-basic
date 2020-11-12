package queue_test

import (
	"fmt"
	"golang-basic/lang/queue"
)

/*
1.godoc事例写法
* func后根据提示直接写测试实例
* 测试实例中写输出：//  output

2.测试文档总结
* 用注释写文档
* 在测试中加入example
* 使用go doc/godoc来查看/生成文档
*/

func ExampleQueue_Pop() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(2)
	q.Push(2)
	q.Push(2)
	q.Push(2)
	q.Pop()

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsNil())

	//		output
	//		2
}
