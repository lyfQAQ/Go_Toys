package tree

// 一个目录下只能有一个包
// 首字母大写 为public  小写为 private
import (
	"testing"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 结构的方法 定义接收者
func (node TreeNode) Print() {
	print(node.value)
}

// 接收者为指针时 调用该函数可以改变内容
func (node *TreeNode) SetValue(value int) {
	node.value = value
}

func TestTree(t *testing.T) {
	var root TreeNode

	root = TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)

	// root.Print()
	// root.left.Print()
	// root.left.SetValue(100)
	// root.left.Print()
	root.InOrder()
}
