package LeetCodeO

import (
	. "GoAlgorithm"
	"strconv"
	"strings"
)

// TODO: NOT COMPLETE
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	strA := serializeTree(A)
	strB := serializeTree(B)
	return strings.Contains(strA, strB)
}

func serializeTree(root *TreeNode) string {
	var res string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			res += strconv.Itoa(node.Val)
		} else {
			res += "#"
		}
		res += " "
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
