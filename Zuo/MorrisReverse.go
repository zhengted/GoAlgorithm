package Zuo

import (
	. "GoAlgorithm/Common"
	"fmt"
	"strconv"
)

func MorrisReserve(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var mostRight *TreeNode

	for cur != nil {
		// 1.判断有没有左树
		mostRight = cur.Left
		if mostRight != nil {
			// 找到cur左树上，真实的最右边节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			// 从while中出来，当前的mostRight一定是cur左树上最右节点
			if mostRight.Right == nil {
				//mostRight == nil
				mostRight.Right = cur
				// fmt.Print(strconv.Itoa(cur.Val)+" ") // 先序遍历插入点
				cur = cur.Left
				continue
			} else {
				//mostRight.Right == cur
				mostRight.Right = nil
			}
		} else {
			//fmt.Print(strconv.Itoa(cur.Val)+" ") // 先序遍历插入点
		}
		//fmt.Print(strconv.Itoa(cur.Val)+" ") // 中序遍历插入点
		cur = cur.Right
	}
	fmt.Println()
}

func MorrisPre(root *TreeNode) {
	if nil == root {
		return
	}
	cur := root
	var mostRight *TreeNode
	for cur != nil {
		if cur.Left != nil {
			mostRight = cur.Left
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				fmt.Print(strconv.Itoa(cur.Val) + " ")
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			fmt.Print(strconv.Itoa(cur.Val) + " ")
		}
		cur = cur.Right
	}
	fmt.Println()
}

func MorrisIn(root *TreeNode) {
	if nil == root {
		return
	}
	cur := root
	var mostRight *TreeNode
	for cur != nil {
		if cur.Left != nil {
			mostRight = cur.Left
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}
		fmt.Print(strconv.Itoa(cur.Val) + " ")
		cur = cur.Right
	}
	fmt.Println()
}

func MorrisPost(root *TreeNode) {
	if nil == root {
		return
	}
	cur := root
	var mostRight *TreeNode
	for nil != cur {
		if cur.Left != nil {
			mostRight = cur.Left
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				edgePrint(cur.Left)
				mostRight.Right = nil
			}
		}
		cur = cur.Right
	}
	edgePrint(root)
}

// 逆序遍历打印右边所有分支
func edgePrint(node *TreeNode) {
	reverseNode(node)
	tmp := node
	for tmp != nil {
		fmt.Print(strconv.Itoa(tmp.Val) + " ")
		tmp = tmp.Right
	}
	reverseNode(node)
}

// 反转链表
func reverseNode(node *TreeNode) {
	var (
		tmp     = node
		worker  *TreeNode
		newHead = &TreeNode{
			0, nil, nil,
		}
	)
	for tmp != nil {
		worker = tmp.Right
		tmpNext := newHead.Right
		tmp.Right = tmpNext
		newHead.Right = tmp
		tmp = worker
	}
}
