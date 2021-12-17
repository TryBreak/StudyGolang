package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node treeNode) print() { //  和 print(node treeNode) 没有区别
	fmt.Printf("%v \n", node.value)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) setValue(value int) { //指针修改
	// (*node).value = value
	if node == nil {
		fmt.Println("node value is nil")
		return
	}
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value} //返回地址
}

func main() {
	var root treeNode
	root = treeNode{
		value: 3,
	}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }

	// fmt.Println(root)
	/*
		root.right.left.setValue(4)
		root.right.left.print()

		root.setValue(100)
		root.print()

		var pRoot *treeNode
		pRoot.setValue(200)
		pRoot = &root
		pRoot.setValue(300)
		pRoot.print()

	*/

	// 遍历树
	root.traverse()

}
