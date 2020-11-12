package infra

import (
	"io/ioutil"
	"net/http"
)

/*
模拟开发组的功能模块
*/

type Retriever struct{}

//注：此处的函数接收者只有类型：省略掉了函数接收者的变量名，因为用不到
func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
