package testing

/*
模拟测试团队测试
*/
type Retriever struct {
}

func (Retriever) Get(url string) string {
	return "fake content"
}
