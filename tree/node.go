package tree

import "fmt"

type Node struct {
	value int
	left, right *Node
}
/*
	go 语言没有构造函数,但是可以用工厂函数
工厂函数就是普通的函数
*/
func CreateNode(value int) *Node {
	return &Node{value: value}
}

// 方法
func (node Node) Print() {
	fmt.Println(node.value)
}

// set 方法 只有指针接收者才能改变结构的内容
func (node *Node) SetValue(value int) {
	node.value = value
}

//// 遍历树
//func (node *TreeNode) Traverse() {
//	if node == nil {
//		return
//	}
//	node.left.Print()
//	node.Print()
//	node.right.Print()
//}

