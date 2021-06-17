package Zuo

import (
	. "GoAlgorithm/Common"
	"fmt"
	"strconv"
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

// 二叉树序列化和反序列化
func SerializeTreePre(root *TreeNode) []string {
	res := []string{}
	serializeTreePre(root, &res)
	return res
}

func serializeTreePre(node *TreeNode, ret *[]string) {
	if node == nil {
		*ret = append(*ret, "#")
		return
	}
	*ret = append(*ret, strconv.Itoa(node.Val))
	serializeTreePre(node.Left, ret)
	serializeTreePre(node.Right, ret)
}

func UnSerializeTreePre(arr *[]string) *TreeNode {
	curValStr := (*arr)[0]
	*arr = (*arr)[1:]
	if curValStr == "#" {
		return nil
	}
	curVal, _ := strconv.Atoi(curValStr)
	ret := &TreeNode{
		Left:  UnSerializeTreePre(arr),
		Right: UnSerializeTreePre(arr),
		Val:   curVal,
	}
	return ret
}

func SerializeTreeByLayer(root *TreeNode) []string {
	ans := []string{}
	queue := []*TreeNode{}
	if root == nil {
		return ans
	}
	ans = append(ans, strconv.Itoa(root.Val))
	queue = append(queue, root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode.Left != nil {
			ans = append(ans, strconv.Itoa(curNode.Left.Val))
			queue = append(queue, curNode.Left)
		} else {
			ans = append(ans, "#")
		}

		if curNode.Right != nil {
			ans = append(ans, strconv.Itoa(curNode.Right.Val))
			queue = append(queue, curNode.Right)
		} else {
			ans = append(ans, "#")
		}
	}
	return ans
}

func UnSerialTreeByLayer(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	curValStr := arr[0]
	arr = arr[1:]
	if curValStr == "#" {
		return nil
	}
	curVal, _ := strconv.Atoi(curValStr)
	root := &TreeNode{
		curVal, nil, nil,
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		curValStr = arr[0]
		arr = arr[1:]
		if curValStr == "#" {
			curNode.Left = nil
		} else {
			curVal, _ := strconv.Atoi(curValStr)
			curNode.Left = &TreeNode{curVal, nil, nil}
			queue = append(queue, curNode.Left)
		}

		curValStr = arr[0]
		arr = arr[1:]
		if curValStr == "#" {
			curNode.Right = nil
		} else {
			curVal, _ := strconv.Atoi(curValStr)
			curNode.Right = &TreeNode{curVal, nil, nil}
			queue = append(queue, curNode.Right)
		}

	}
	return root
}
