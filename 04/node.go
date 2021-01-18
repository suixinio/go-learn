package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	// 使用自定义工厂函数
	// 注意返回了局部变量的地址
	return &treeNode{value: value}
}

func (node treeNode) setValue(value int) {
	// 值传递
	node.value = value
}

//func (node *treeNode) setValue(value int) {
//	// 指针传递
//	node.value = value
//}

func (node treeNode) print() {
	// 前面的 (node treeNode) 是接受者
	fmt.Println(node.value)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	fmt.Println(root)
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.print()
	fmt.Println("----------")
	root.traverse()
}
