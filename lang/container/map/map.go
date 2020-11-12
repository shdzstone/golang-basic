package main

import "fmt"

/**
1.map及复合map定义
* m:=map[k]v,m:=map[k1]map[k2]v
* m:=make(map[string]int)
* var m map[string]int
* map中的键值对是无序的，是hashmap

2.获取元素
* m[key]
* key不存在时，获得value类型的初始值：zero value
* 用value,ok:=m[key]来判断是否存在key
* 用delete删除一个key

3.map遍历
* 使用range遍历key,或者key,value对
* 不保证遍历顺序，如需顺需，需手动对key排序：可以把key加到某slice中，slice可以排序
* 使用len()来获取元素个数

4.map的key
* map使用哈希表，必须可以比较相等
* 除了slice,map,function的内建类型都可以做为key
* 自定义类型Struct若不包含上述字段，也可以做为key，编译器在编译时会检查

*/
func main() {
	fmt.Println("Creating map")
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "go",
		"site":    "imocc",
		"quality": "notbad",
	}
	//go语言中的nil和empty都是可以参与运算的，和其它语言不一样
	m2 := make(map[string]int) //m2==empty map
	var m3 map[string]int      //m3==nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traveling map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//map取不存在的key时，会输出zero value
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting map")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
