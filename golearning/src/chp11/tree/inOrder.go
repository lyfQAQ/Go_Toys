package tree

func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	node.Print()
	node.right.InOrder()
}
