package main

import (
	"fmt"
	"golang-basic/lang/tree"
)

/**
1. 包
* 每个目录只能有一个包
* 包名不一定和目录名一样
* main包包含可执行入口main函数，目录下若有一个main函数，则该目录下只能有一个main包
* main函数不在main包内时，运行不了
* 为结构定义的方法必须放在同一个包内
* 可以是不同的文件
* 每个package下只能有一个main函数,否则项目整体编译时会报错：main redeclared in this block

2.包的封装
* 包括变量名、类型名、函数名
* 封装特性指的是包，而不是同一文件
* 名字一般使用CamelCase，不使用"_"开头或数字
* 首字母大写：public
* 首字母小写：private

3.包的引用
* 用法：package_name.method_name[struct_name][interface_name]

4. 如何扩充系统类型或者别人的类型
* 定义别名：最简单，但转换麻烦
* 使用组合：最常用
* 使用内嵌：需要省下许多代码，要求go语言基础好

*/

//使用组合扩充tree.Node
//有一个tree.Node的指针成员
//组合不一定放指针，也可以放对象
//组合，可以理解为定义一个结构体中，其字段可以是其他的结构体，这样，不同的结构体就可以共用相同的字段
//注意，结构体不能包含自身，但可能包含指向自身的结构体指针
type myTreeNode struct {
	node *tree.Node
}

//后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	leftNode := myTreeNode{myNode.node.Left}
	leftNode.postOrder()
	rightNode := myTreeNode{myNode.node.Right}
	rightNode.postOrder()
	myNode.node.Print()
}

func main() {
	//声明一个empty实例
	//var root treeNode
	root := tree.Node{}

	fmt.Println(root)

	root = tree.Node{Value: 3}
	//"."操作符取成员实例
	root.Left = &tree.Node{}

	//go语言中，"."操作符可以直接取指针成员，而不是->
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	//使用new关键字来创建empty实例
	root.Right.Left = new(tree.Node)
	fmt.Println(root)

	/*
		//使用:=初始化对象数组
		nodes := []treeNode{
			{value: 3},
			{},
			{6, nil, &root},
		}
		fmt.Println(nodes)
	*/

	root.Print()
	fmt.Println()

	root.Right.Left.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Left.Right = tree.CreateNode(2)

	/*
		var pRoot *treeNode
		//空指针进行值传递时，会报panic
		//pRoot.print()
		pRoot.setValue(200)
		pRoot = &root
		pRoot.setValue(300)
		pRoot.print()
	*/
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	//使用闭包来数结点的数量
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	//使用channel遍历树
	maxValue := 0
	ch := root.TraverseWithChannel()
	for node := range ch {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxValue)
}
