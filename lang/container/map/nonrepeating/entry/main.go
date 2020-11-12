package main

import (
	"fmt"
	"golang-basic/lang/container/map/nonrepeating"
)

func main() {
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("aadfasdiopaipsiupaisofu"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("asd"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("090909090"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("a"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("123"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("abcdefg"))
	fmt.Println(nonrepeating.LengthOfNonRepeatingSubStr("一二三四三二一"))
}
