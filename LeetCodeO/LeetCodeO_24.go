package LeetCodeO

import . "GoAlgorithm/Common"

func reverseList(head *ListNode) *ListNode {
	// 1. 头插法
	if nil == head {
		return nil
	}
	newHead := new(ListNode)
	for head != nil {
		temp := &ListNode{
			Val:  head.Val,
			Next: newHead.Next,
		}
		newHead.Next = temp
		head = head.Next
	}
	return newHead.Next
}
