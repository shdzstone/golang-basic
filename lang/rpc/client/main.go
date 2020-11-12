package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "golang-basic/lang/rpc"
)

/*
1.远程调用客户端
* 使用net.Dial()建立到rpc服务器的tcp连接
* 以上面建立的tcp连接为参数建立jsonrpc.NewClient()
* 使用client.Call()调用rpc服务的相应函数
* Call()函数签名：func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
* 注：reply一般是个指针，用以接收远程调用处理的结果
*/

func main() {
	conn, err := net.Dial("tcp", ":4321")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	//使用call()来包装rpc服务的函数，并调用
	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
