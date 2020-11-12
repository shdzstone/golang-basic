package main

import (
	"fmt"
	"log"
	"regexp"
)

/*
1.正则表达式
* 用以获取匹配格式的字符串
* 正则表达式有自己的语法
* regexp.Compile()编译并解析正则表达式的语法，若编译通过，则执行并返回执行结果，否则error
* regexp.MustCompile()直接执行，不编译解析语法，不符合语法则panic
* re.FindString()查找的字符串
* ``里面的字符串不发生转义
* re.FindAllStringSubmatch()
*/

const text = `my email is ccmouse@gmail.com@abc
email1 is abc@def.org
email2 is kkkk@qq.com
email3 is ddd@abc.com.cn
`

var idUrlRe = regexp.MustCompile(`http://album.parser.com/u/([0-9]+)`)

func extractString(contents []byte, re *regexp.Regexp) string {
	log.Printf("url:%s", contents)
	match := re.FindSubmatch(contents)

	for _, item := range match {
		log.Printf("match item:%s", item)
	}

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func main() {
	//返回值是*Regexp
	//正则表达式提取功能：加()表示提取出()内的结果
	//re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	//返回结果是[][]string二维slice
	//match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)
	//for _, val := range match {
	//	fmt.Println(val)
	//}
	url := "http://album.parser.com/u/1933829515"
	s := extractString([]byte(url), idUrlRe)
	fmt.Println(s)
}
