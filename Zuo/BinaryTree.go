package Zuo

import (
	. "GoAlgorithm/Common"
	"fmt"
)

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PrintTree(root.Left)
	PrintTree(root.Right)
}

type TreeNodeWithPar struct {
	val    int
	left   *TreeNodeWithPar
	right  *TreeNodeWithPar
	parent *TreeNodeWithPar
}

// 寻找结点的中序后继
func FindPostNode(node *TreeNodeWithPar) *TreeNodeWithPar {
	if node == nil {
		return nil
	}
	// 有右孩子 那后继是右孩子的最左节点
	if node.right != nil {
		node = node.right
		for node.left != nil {
			node = node.left
		}
		return node
	}
	// 无右孩子 向上找到cur.left == 当前节点的祖宗  返回cur
	parent := node.parent
	for parent != nil && parent.right == node {
		node = parent
		parent = node.parent
	}
	return parent

}

// 折痕问题
func PrintFold(N int) {
	printProcess(1, N, true)
}

func printProcess(i, N int, down bool) {
	// i：层数 N：固定值 表示总层数 down:是否为凹折痕
	if i > N {
		return
	}
	printProcess(i+1, N, true)
	var content string
	if down {
		content = "down "
	} else {
		content = "up"
	}
	fmt.Print(content)
	printProcess(i+1, N, false)
}
