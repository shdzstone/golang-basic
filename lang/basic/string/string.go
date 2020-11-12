package main

import (
	"fmt"
	"unicode/utf8"
)

/**
1.rune与byte定义
* 计算机是二进制的，字符最终也是转换成二进制保存起来的。
* 字符集就是定义字符对应的数值。
* ASCⅡ是一种单字节字符集。
* unicode也是一个字符集，为每个字符规定一个用来表示该字符的数字，但是并没有规定该数字的二进制保存方式。
* utf-8规定了对于unicode值的二进制保存方式，字符的unicode值决定了字符需要用多少字节表示
* utf-8是一种变长编码：英文字符1个字节，阿拉伯/拉丁文字符2个字节，中文字符3个字节，古希腊文4个字节。

//byte是uint8的别名，表示1个字节，在所有方面都等同于uint8。按照惯例，它用于区分字节值和8位无符号整数值。代表一个ASCII字符。
type byte = uint8
//rune是int32的别名，表示4个字节，在所有方面都等同于int32。按照惯例，它用于区分字符值和整数值。代表一个UTF-8字符。
type rune = int32

2.字符串定义
type string struct {
    data uintptr
    len int
}
type slice struct {
    data uintptr
    len int
    cap int
}
* go语言中，字符串默认使用utf-8编码。
* Go语言中，string就是只读的采用utf8编码的字节切片(slice)。
* 因此，用len函数获取到string的长度并不是字符个数，而是字节个数。 for循环遍历[]byte时输出的也是各个字节。

3. rune使用
* 使用range遍历pos,rune对，若含有中文字符，这里的pos是不连续的
* utf8.DecodeRune([]byte)获取字节数组中的字符
* utf8.RuneCountInString获取字符数量
* len(string)获取的是字符串的字节长度
* 将string转换成[]byte，转换后为字节码
* 将[]byte转换成[]rune，转换后为字符切片

4. 字符串操作
* 内置的strings包中有字符串操作，strings.Join()、strings.Map()、strings.Replace()、strings.NewReplacer()等
* Fields、Split、Join字符串拆分或组合
* Contains、Index查找字符串
* ToLower、ToUpper
* Trim、TrimRight、TrimLeft
*/

func main() {

	s := "go社区!" //utf-8编码

	//len counts the number of bytes in a string
	fmt.Println(len(s))
	fmt.Printf("%s\n", []byte(s))
	//individual characters
	for _, v := range []byte(s) {
		fmt.Printf("%c ", v)
	}
	fmt.Println()

	//使用range遍历字符串时，每个中文字符算三个字节长度
	for i, ch := range s { //ch is a rune,ch类型为int32，实际上是一个4字节的整数
		fmt.Printf("(%d %X)  ", i, ch)
	}
	fmt.Println()
	fmt.Printf("Rune count:%d\n", utf8.RuneCountInString(s))

	//转换为[]byte后，使用utf8解析为rune
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		//取字符，utf-8转unicode的一个过程
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	// 将string转换为[]rune后，返回转换后的字符切片，每个中文字符算一个长度
	for i, ch := range []rune(s) {
		//取16进制数值，即为该字符的unicode编码值
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
}
