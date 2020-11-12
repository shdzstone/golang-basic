package main

import (
	"golang-basic/lang/errorhandling/filelistingserver/errwrapper"
	"golang-basic/lang/errorhandling/filelistingserver/filelisting"
	"net/http"
	_ "net/http/pprof"
)

/*
1.http标准库使用
* http.HandleFunc()初始化Http请求的路由及Http request/response处理函数
* http.ListenAndServe()监听tcp端口并启动http服务器

2.http服务器的性能分析
* import _ "net/http/pprof"
* 访问/debug/pprof/
* 使用go tool pprof分析性能
* go tool pprof http://localhost:8888/debug/pprof/profile
* go tool pprof http://localhost:8888/debug/pprof/heap
*/

func main() {
	uri := "/list/"
	http.HandleFunc(uri, errwrapper.ErrWrapper(filelisting.FileHandlerList))
	http.ListenAndServe(":8888", nil)

}
