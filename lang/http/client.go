package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

/*
1.http
* http.Get()：使用默认的http client发送请求
* http.NewRequest()：创建http请求
* http.Client：控制请求头部
* 使用httputil.DumpResponse简化工作

2.其他标准库
* bufio:可以一次读写很多数据，再一次flush到硬盘
* log:打印日志
* encoding/json：序列化及反序列化
* regexp：正则表达式
* time:time.Duration、time.Sleep、time.After、time.Tick、time.second
* time:产生很多chan，与select配合，对超时、定时任务进行处理
* strings/math/rand

3. godoc查看项目文档及其他标准库的文档
* godoc -http :8888
* https://studygolang.com/pkgdoc
*/
func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	//request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", s)

}
