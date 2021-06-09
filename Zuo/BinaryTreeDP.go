package Zuo

import (
	. "GoAlgorithm/Common"
	"math"
)

//二叉树递归套路
//1.假设以X节点为头，假设可以向X左树X右树要任何信息
//2.在上一步假设下，讨论以X为头节点的树，得到答案的可能性
//3.列出所有可能性后，确定到底需要左树右树什么信息
//4.把左树信息和右树信息求全集，就是任何一颗子树需要返回的信息S
//5.递归函数都返回S，每一棵子树都这么要求
//6.写代码，在代码中考虑如何把左树信息和右树信息整合出整棵树的信息

func IsBalanceTree(root *TreeNode) bool {
	// 1.左树是平衡的 2. 右树是平衡的 3.左树和右树的高度差不超过1
	ret, _ := getBalanceAndHeight(root)
	return ret
}

func getBalanceAndHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lBal, lHeight := getBalanceAndHeight(root.Left)
	rBal, rHeight := getBalanceAndHeight(root.Right)
	curHeight := intMax(lHeight, rHeight) + 1
	return lBal && rBal && math.Abs(float64(lHeight)-float64(rHeight)) < 2, curHeight
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 返回二叉树两个节点的最大距离
func GetTreeMaxDistance(root *TreeNode) int {
	// 1.最大距离与X无关（不经过X）  return maxInt(leftDistance,rightDistance)
	// 2.最大距离通过X return leftHeight+rightHeight+1
	ret, _ := getTreeDistanceAndHeight(root)
	return ret
}

func getTreeDistanceAndHeight(node *TreeNode) (maxDistance, height int) {
	if node == nil {
		return 0, 0
	}
	lDis, lHeight := getTreeDistanceAndHeight(node.Left)
	rDis, rHeight := getTreeDistanceAndHeight(node.Right)
	height = intMax(lHeight, rHeight) + 1
	maxDistance = intMax(intMax(lDis, rDis), height)
	return maxDistance, height
}

type employee struct {
	Happy int
	next  []employee
}

type maxHappyInfo struct {
	yes int // 来的情况下该节点的最大快乐值
	no  int // 不来的情况下该节点的最大快乐值
}

// 员工派对的最大快乐值
func MaxHappy(root *employee) *maxHappyInfo {
	if len(root.next) == 0 {
		return &maxHappyInfo{
			yes: root.Happy,
			no:  0,
		}
	}
	ret := &maxHappyInfo{}
	ret.yes = root.Happy
	ret.no = 0
	for _, e := range root.next {
		info := MaxHappy(&e)
		ret.yes += info.no
		ret.no += intMax(info.yes, info.no)
	}
	return ret
}
