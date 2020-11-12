package mock

import "fmt"

/*
* option/alt+enter:快速实现接口快捷键
* Retriever并没有显式的声明实现了某个接口，只是通过编译器"实现"某个接口的方，
* 所以，只要实现了"接口"的方法，即是实现了该接口
* 从go语言的角度来讲，并不需要显式的"说明"：实现了某接口
 */
type Retriever struct {
	Contents string
}

func (r *Retriever) Get(s string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

//stringer interface
func (r Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", r.Contents)
}
