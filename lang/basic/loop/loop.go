package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
1.for循环
* 没有while
* for的条件里不需要括号
* for的条件里可以省略初始条件，结束条件，递增表达式
* for省略初始条件,相当于while
* for初始条件和递增条件都省略掉,相当于while
* for省略所有条件，死循环
*/

//for基本用法
func sumLoop() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

//二进制转换
//for省略初始条件,相当于while
func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		//字符串相加即字符串拼接
		//int转字符串用strconv.Itoa
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//打开文件
//for初始条件和递增条件都省略掉,相当于while
func printFile(filename string) {
	//打开文件，返回一个file对象、err对象
	file, err := os.Open(filename)
	if err != nil {
		//程序停下来，报错
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	//bufio.NewScanner一行一行的扫描内容
	//file实现了Reader接口，因此参数是io.Reader接口，而不是file对象
	scanner := bufio.NewScanner(reader)
	//省略初始条件和递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//省略所有条件，死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//参数是writer，只要实现writer接口的file、http连接均可使用
	//fmt.Fprintf()

	//参数是reader，只要实现reader接口的file、http连接即可
	//fmt.Fscanf()

	fmt.Println(
		sumLoop(),
		convertToBinary(5),  //101
		convertToBinary(13), //1101
		convertToBinary(234234234),
	)
	printFile("varOp/abc.txt")
	//forever()

	str := `abc"d"
			kkk
			123123
	
			adsfadsf`
	//使用file文件的地方换成reader之后，使用范围就广了，可以使用许多系统库
	printFileContents(strings.NewReader(str))
}
