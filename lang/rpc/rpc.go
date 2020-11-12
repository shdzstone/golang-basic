package rpcdemo

import "errors"

/*
1.rpc服务声明
* 需要使用package定义一些数据结构来支持rpc
* rpc调用方式：Service.Method
* go语言rpc函数参数要求：args,result，返回值：error

2.rpc模块包装成一个server，对外提供服务
* 向rpc注册服务
* 开启rpc服务，监听端口
* 循环接收端口传来的connection连接
* 使用jsonrpc.ServeConn()处理接收到的连接

3.rpc调用通过rpc client调用
* 使用net.Dial()建立到rpc服务器的tcp连接
* 以上面建立的tcp连接为参数建立jsonrpc.NewClient()
* 使用client.Call()调用rpc服务的相应函数
* Call()函数签名：func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
* 注：reply一般是个指针，用以接收远程调用处理的结果

*/

type DemoService struct{}

type Args struct {
	A, B int
}

//RPC方法的声明规范：func (RPCService) MethodName(args interface{}, result interface{}) error
//result必须为指针，以便客户调用结束后将结果写进接收参数
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
