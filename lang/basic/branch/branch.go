package main

import (
	"fmt"
	"io/ioutil"
)

/**
1.条件语句
* if的条件里不需要括号
* if的条件里可以赋值：if contents, err := ioutil.ReadFile(filename); err != nil {}
* if条件赋值的变量作用域就在这个if语句里

2.switch
* switch会自动break,除非使用fallthrough
* switch后可以没有表达式，在case中加入条件即可
* switch不需要break,也可以直接switch多个条件
*/

//if
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

//多返回值用法，一般是普通返回值+error
//函数error返回值方法
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation:" + op)
	}
}

//读取文件、if
func fileReader() {
	const filename = "lang/varOp/abc.txt"
	//一般用法
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//简略用法
	//注：此处的contents,err只有在条件内使用，外部不能引用
	//ioutil.ReadFile读取文件中所有内容，按行组成数组
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//switch不跟表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		//panic为了报错
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
		//default:
		//	panic(fmt.Sprintf("Wrong score:%d", score))
	}
	return g
}

func main() {
	fmt.Println(bounded(101))
	fileReader()
	fmt.Println(eval(13, 13, "+"))

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
	)
}
