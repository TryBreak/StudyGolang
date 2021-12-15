
type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node treeNode) print() { // 给对象放置方法
	fmt.Println("墨七", node)
}

func main() {
	var root treeNode
	root = treeNode{3, nil, nil}
	root.left = &treeNode{1, nil, nil}
	root.right = &treeNode{2, nil, nil}
	root.left.left = &treeNode{3, nil, nil}
	root.left.right = &treeNode{4, nil, nil}
	root.right.left = &treeNode{5, nil, nil}
	root.right.right = &treeNode{6, nil, nil}
	root.print()
	fmt.Println(root)

	fmt.Println("----root2--------------")
	var root2 treeNode
	root2.print()
	fmt.Println(root2)
}