package main

import (
	"fmt"
	"golang-basic/lang/interface/retriever/mock"
	"golang-basic/lang/interface/retriever/real"
	"time"
)

/*
1.接口类型的定义
#go语言中，所有类型的内部实现
type _type struct {
	size       uintptr // type size
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldalign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	alg        *typeAlg  // algorithm table
	gcdata    *byte    // garbage collection data
	str       nameOff  // string form
	ptrToThis typeOff  // type for pointer to this type, may be zero
}
#空接口
#empty interface
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
#非空接口
#if non-empty interface
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
#接口和实现者类型的混合结构体
type itab struct {
	inter  *interfacetype
	_type  *_type
	link   *itab
	bad    int32
	inhash int32      // has this itab been added to hash?
	fun    [1]uintptr // variable sized
}
#接口类型的内部实现结构体
type interfacetype struct {
	type    _type
	pkgpath name
	mhdr    []imethod
}

2. 接口变量的内部结构
* eface:pair<*_type, unsafe.Pointer>
* iface:pair<*itab, unsafe.Pointer>
* itab:pair<interface type, concrete type>
* 接口变量是一个结构体，内部有两个指针：itab/_type指针和data指针，itab指针为<interface type, concrete type>结构体
* 接口变量赋值时，会触发类型检查，编译器会将接口变量的data指向rvalue或rvalue的拷贝；
* 接口变量赋值时，rvalue可以是一个变量，编译器将拷贝后的值的内存地址赋值给接口变量的data指针；
* 接口变量赋值时，rvalue也可以是一个指针，编译器将rvalue指针赋值给接口变量的data指针；
* 接口变量赋值时，实现者的值还是指针，需要看具体实现者的函数接收者是指针还是变量
* 可以直接打印接口变量的值，，也可以通过type switch判断其类型，还可以通过Type assertion来判断接口变量的类型

3.接口变量的实现
* 只要实现接口里的方法,即实现了该接口：接口的实现是隐式的
* 接口实现者的函数接收者类型，可以是实现者的对象，也可以是实现指的指针
* 要避免使用接口指针，因为接口变量结构体内本身就含有接口自身和实现者的指针
* 接口变量赋值同样采用值传递
* 若接口实现为指针接收者，接口变量赋值时rvalue必须为指针
* 若接口实现为值接收者，接口变量赋值时rvalue既可以为指针，也可以为值

4. 扩展
* go语言所有类型默认都实现了interface接口
* interface{}可以用来表示任何类型

5. 接口组合
* 接口组合是对使用者来说的，使用者可以自己定义并实现接口中的方法
* 使用者也可以把现有的接口组合起来去定义一个新的接口
* 使用者无论怎样定义，对于实现者来讲，只需要所要的实现方法即可，可以做到代码的解耦
* go语言的接口赋予了编程很的灵活性，允许在使用的时候去描述一个接口具有怎样的功能，比如：io.ReadWriter()、io.ReadCloser()、io.ReadWriteCloser()

6. 常用的系统接口
* stringer接口
* io.Reader、io.Writer
* go语言中，底层相关的读写做成Writer/Reader后，可以使用系统库,比如Fprintf、Fscanf

7. 接口变量
* 表示任何类型：interface{}
* type assertion：类型转换
* switch assertion：类型判断
*/

//Retriever接口
type Retriever interface {
	Get(string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

//Poster接口
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

//接口组合
type RetrieverPoster interface {
	Poster
	Retriever
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "Another fake imooc",
	})
	return s.Get(url)
}

//接口的值:
//1.实现者的类型：type类型或指针类型
//2.实现者的值或指针，可以通过r.(type)断言将接口值转换为相应的receiver
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T  %v  \n", r, r)
	fmt.Printf("> Type switch:")
	//switch assertion
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Printf("Contents:%v\n", v)
	case *real.Retriever:
		fmt.Printf("UserAgent:%v\n", v.UserAgent)
	}
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{Contents: "this is a fake imooc.com"}
	r = retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "mozilla/5.0",
		TimeOut:   time.Second,
	}
	inspect(r)

	//Type assertion:通过.(type)将接口类型转换为receiver的实际类型
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever")
	}
	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}
