package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcdemo "golang-basic/lang/rpc"
)

/*
1.rpc服务开启流程
* 向rpc注册服务
* 开启rpc服务，监听端口
* 循环接收端口传来的connection连接
* 使用jsonrpc.ServeConn()处理接收到的连接

2.telnet测试jsonrpc服务
* telnet URL port
* 传输jsonrpc可以接收的json调用，格式为:{"method":"abc.def"}

*/

func main() {
	//首先注册服务
	rpc.Register(rpcdemo.DemoService{})

	//开启server:监听TCP的1234端口
	listener, err := net.Listen("tcp", ":4321")
	if err != nil {
		panic(err)
	}

	for {
		//接收传入的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error:%v", err.Error())
			continue
		}
		fmt.Println("receive connection")
		//使用rpc当场处理掉conn，开启goroutine是为了不让该处理过程阻塞当前goroutine
		go jsonrpc.ServeConn(conn)
	}
}
