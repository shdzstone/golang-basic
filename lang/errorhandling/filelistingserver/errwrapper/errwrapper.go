package errwrapper

import (
	"fmt"
	"net/http"
	"os"
)

/*
1. go语言中错误处理
* defer + panic + recover
* type assertion
* 把业务处理和错误处理分开
* 有些内部错误不能被用户看到，但有些用户输入错误可以让用户看到

2.error vs panic
* panic未使用recover捕获的结果是程序直接挂掉，error则可以通过代码捕获解决
* 什么时候使用error，什么时候使用panic：尽量使用error，尽量不要使用panic
* 意料之中的：使用error。如：文件打不开
* 意料之外的：使用panic。如：数组越界

3. 函数式编程
* 函数又能做参数，又能做返回值
* 函数式编程：输入参数和输出参数均为函数，把输入的函数包装下，包装成一个输出的函数来输出
*/

type AppHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

//注：接收参数和返回值均是函数
// 函数式编程：输入参数和输出参数均为函数，把输入的函数包装下，包装成一个输出的函数来输出
func ErrWrapper(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//defer...recover
		defer func() {
			if r := recover(); r != nil {
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			fmt.Printf("Error occurred"+" handling request: %s\n", err.Error())

			//自定义user error：可以让用户看到的错误类型
			//type assertion
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			//switch assertion
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			//http.StatusText用户报错，以免暴露内部报错信息
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
