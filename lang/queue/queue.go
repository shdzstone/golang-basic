package queue

/**
1.使用别名的方式来扩充系统类型或别人的类型
* 定义别名：最简单，但转换成使用组合比较麻烦，改动较大
* interface{}可以用以表示任何类型

2.go生成文档及示例代码

* go doc直接在命令行查看文档
* go doc Queue 查看Queue具体的文档，包含具体的方法及注释
* go help doc 查看具体go doc 命令
* go doc fmt.Println 查看系统库fmt

3.godoc
* godoc在go1.13被拿了，需要重新安装
* go get golang.org/x/tools/cmd/godoc
* godoc查看具体命令
* godoc -http:6060 可在浏览器中查看本机转换后的文档
* godoc中紧挨着代码行的单行注释可以显示在文档中
* godoc单行注释中空两格可以显式例子，例如：//		e.g. q.Push(123)
*/

//A FIFO queue
type Queue []interface{}

//函数接收者为指针
//故可以直接改变q的值

//Push thd elements into the queue
//		e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsNil() bool {
	return len(*q) == 0
}
