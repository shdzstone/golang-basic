package main

import (
	"fmt"
	"golang-basic/lang/interface/interface/infra"
)

/**
1.接口概念
* go工程化管理，模块化、可配置、可测试
* 强类型语言：熟悉接口的概念
* 弱类型语言：没(少)有接口的概念
* 小孩才分对错
* 大人只看利弊
* 接口是一个概念，只是讲：只要能调用某个方法即可

2.duck typing
* 大黄鸭是鸭子吗
* 传统类型系统：脊椎动物门、脊椎动物亚门、鸟纲雁形目...
* duck typing：是鸭子，"像鸭子走路，像鸭子叫（长的像鸭子），那么就是鸭子"
* 接口描述事物的外部形为而非内部结构
* 从使用者角度来看：对于孩子来讲，是鸭子，对于吃货来讲，不是鸭子
* 严格说go属于结构化类型系统，类似duck typing
* duck typing是动态绑定，但go语言是编译时就绑定的，死扣动态绑定的话，go语言不是duck typing
* 但从描述事物的外部形为而非内部结构的特点来看，go语言的接口，也是duck typing

3. python中的duck typing
def download(retriever):
	return retriever.get("https://www.imooc.com")
* 运行时才知道传入的retriever有没有get
* 需要注释来说明接口

4. c++中的duck typing
template <class R>
string download(const R& retriever) {
	return retriever.get("https://www.imooc.com");
}
* 编译时才知道传入的retriever有没有get
* 需要注释来说明接口

5. java中的duck typing
<R extends Retriever>
string download(R r){
	return r.get("https://www.imooc.com")
}
* 传入的参数必须是实现retriever接口
* 不是duck typing
* 在编译时就能知道，r是否有get
* 同时需要Readable,Appendable怎么办？(apache polygene)
* 无法直接实现多接口的实现

6. go语言中的duck typing
* 同时需要Readable,Appendable怎么办？(apache polygene)
* 同时具有python、c++的duck typing的灵活性，"只要实现某个方法即可"
* 又具有java的类型检查

7. go语言中接口的定义和实现
* 使用者(download)->实现者(retriever)
* 接口由"使用者"来定义
*/

// ?: Something that can "Get"
type Retriever interface {
	Get(string) string
}

//确定具体接口的实现
//此处，infra.Retriever并没有直接"声明并实现"Retriever，但可以直接拿过来用
func getRetriever() Retriever {
	return infra.Retriever{}
}

func main() {
	var r Retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))

}
