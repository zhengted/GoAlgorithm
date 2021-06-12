package Zuo

import . "GoAlgorithm/Common"

func ReverseList(head *ListNode) *ListNode {
	var (
		pre  *ListNode
		next *ListNode
	)
	for head != nil {
		// 1. 标识下一个节点
		next = head.Next
		// 2. 处理节点指针
		head.Next = pre
		// 3. 往下走
		pre = head
		head = next
	}
	return pre
}

func ReverseDoubleList(head *DoubleNode) *DoubleNode {
	var (
		pre  *DoubleNode
		next *DoubleNode
	)
	for head != nil {
		// 1. 标识下一个节点
		next = head.Next
		// 2. 处理节点指针
		head.Next = pre
		head.Pre = next
		// 3. 往下走
		pre = head
		head = next
	}
	return pre
}

func RemoveNumInList(head *ListNode, num int) *ListNode {
	// 处理头部，找出第一个不需要删的位置
	for head != nil && head.Val == num {
		head = head.Next
	}
	// 记录头部
	cur, pre := head, head
	for cur != nil {
		if cur.Val == num {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
