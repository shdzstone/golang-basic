package main

import (
	"fmt"
	"golang-basic/lang/tree"
)

/**
1.内嵌的方式扩充系统类型或别人的类型
* 内嵌Embedding的方式实际上就是一个语法糖，可以省去很多代码量
* 内嵌时可以省略成员名，但实际上，该成员的名字默认为最后一个字段的名，比如：*tree.Node，最后一个字段为Node，那么该成员的名字默认就为Node
* 内嵌有点相似于继承，但不同
* 内嵌可以重载函数
* 内嵌只是看起来相似于父类和子类的继承，实际上，内嵌对象和被内嵌的对象没有任何关系
* Go语言支持直接将类型作为结构体的字段，而不需要取变量名，这种字段叫匿名字段
* 通过匿名字段组合其他类型，而后访问匿名字段类型所带的方法和字段时，不需要使用叶子属性
*/
type embeddingTreeNode struct {
	*tree.Node //Embedding
}

//后序遍历
func (myNode *embeddingTreeNode) embeddingPostOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	//内嵌可以直接获取其成员下属的成员变量及成员函数
	leftNode := embeddingTreeNode{myNode.Left}
	leftNode.embeddingPostOrder()
	rightNode := embeddingTreeNode{myNode.Right}
	rightNode.embeddingPostOrder()
	myNode.Print()

}

//重载
//go语言中不叫重载，叫shadowed method
//若使用embedding直接调用Traversal，则执行本函数
//若使用embedding中的成员调用Traversal，则执行shadowed函数
func (myNode *embeddingTreeNode) Traverse() {
	fmt.Println("This method is shadowed")
}

func main() {
	//声明一个empty实例
	//var root treeNode
	root := embeddingTreeNode{&tree.Node{Value: 3}}
	//"."操作符取成员实例
	root.Left = &tree.Node{}

	//go语言中，"."操作符可以直接取指针成员，而不是->
	root.Right = &tree.Node{5, nil, nil}
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
	fmt.Println("In-order traversal:")
	fmt.Printf("root.Traverse():")
	root.Traverse()

	fmt.Printf("root.Node.Traversal():")
	root.Node.Traverse()
	fmt.Println()

	fmt.Println("My own post-order traversal:")
	root.embeddingPostOrder()
	fmt.Println()

	//var baseRoot *tree.Node
	//baseRoot = &root
}
