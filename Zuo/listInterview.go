package Zuo

import (
	. "GoAlgorithm/Common"
	"math"
)

// 奇数长返回中点，偶数长返回上中点
func MidOrUpMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	// 链表有三个点或以上
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 返回中点或下中点
func MidOrDownMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	// 链表有三个点或以上
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 判断链表是否为回文结构
//   栈实现 三种方法
// 1. 依次加入栈 弹出和原链表对比
// 2. 快慢指针到中点或上中点，慢指针后续的加入到栈里，头节点遍历与弹出对比
// 3. 快慢指针到中点或上中点，慢指针后面的部分反转链表变为 1->2->3<-2<-1

func IsReverseList1(head *ListNode) bool {
	if head == nil {
		return true
	}
	node := head
	stack := []*ListNode{}
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		if node.Val != head.Val {
			return false
		}
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]
		head = head.Next
	}
	return true
}

func IsReverseList2(head *ListNode) bool {
	if head == nil {
		return true
	}
	var findMid func(head *ListNode) *ListNode
	findMid = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil || head.Next.Next == nil {
			return head
		}
		fast := head.Next.Next
		slow := head.Next
		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}
	mid := findMid(head)
	node := mid
	stack := []*ListNode{}
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}

	for len(stack) > 0 {
		nLen := len(stack)
		node = stack[nLen-1]
		if node.Val != head.Val {
			return false
		}
		stack[nLen-1] = nil
		stack = stack[:nLen-1]
		head = head.Next
	}
	return true
}

func IsReverseList3(head *ListNode) bool {
	if head == nil {
		return true
	}
	var findMid func(head *ListNode) *ListNode
	findMid = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil || head.Next.Next == nil {
			return head
		}
		fast := head.Next.Next
		slow := head.Next
		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}
	mid := findMid(head)
	var reverse func(node *ListNode) *ListNode // 以node.Next节点为头节点反转链表，返回新头部
	reverse = func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}
		pTail := node
		cur := node.Next
		for cur != nil {
			next := cur.Next
			cur.Next = pTail
			pTail = cur
			cur = next
		}
		return pTail
	}
	pR := reverse(mid)
	for pR != nil && head != nil {
		if pR.Val != head.Val {
			return false
		}
		pR = pR.Next
		head = head.Next
	}
	return true
}

// 将单链表转换成左边小 中间相等 右边大的形式
func ListPartition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := []int{}
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	less := -1
	more := len(temp)
	for i := 0; i < more; {
		if temp[i] < x {
			temp[i], temp[less+1] = temp[less+1], temp[i]
			less += 1
			i += 1
		} else if temp[i] > x {
			temp[i], temp[more-1] = temp[more-1], temp[i]
			more -= 1
		} else {
			i++
		}
	}
	newHead := &ListNode{}
	p := newHead
	for i := 0; i < len(temp); i++ {
		p.Val = temp[i]
		if len(temp)-1 == i {
			break
		}
		p.Next = &ListNode{}
		p = p.Next
	}
	return newHead
}

func SwapInt(i, j *int) {
	*i = *i ^ *j
	*j = *i ^ *j
	*i = *i ^ *j
}

// 随机链表复制：复制一个包含随机节点的链表
type RandNode struct {
	val  int
	next *RandNode
	rand *RandNode
}

func CloneRandomList1(head *RandNode) *RandNode {
	if head == nil {
		return nil
	}
	hash := make(map[*RandNode]*RandNode)
	pOld := head
	for pOld != nil {
		hash[pOld] = &RandNode{
			val:  pOld.val,
			next: nil,
			rand: nil,
		}
		pOld = pOld.next
	}
	pOld = head
	for pOld != nil {
		newNode := hash[pOld]
		oldNext := pOld.next
		newNext := hash[oldNext]
		oldRand := pOld.rand
		newRand := hash[oldRand]
		newNode.next = newNext
		newNode.rand = newRand

		pOld = pOld.next
	}

	return hash[head]
}

// 优化方法
// 1-->2-->3  ==>  1-->1'-->2-->2'-->3-->3'
func CloneRandomList2(head *RandNode) *RandNode {
	if head == nil {
		return nil
	}
	cur := head
	next := &RandNode{}
	for cur != nil {
		next = cur.next
		cur.next = &RandNode{cur.val, next, nil}
		cur = next
	}
	cur = head
	curCopy := cur.next
	for cur != nil {
		next = cur.next.next
		curCopy = cur.next
		if cur.rand != nil {
			curCopy.rand = cur.rand.next
		} else {
			curCopy.rand = nil
		}
		cur = next
	}
	ret := head.next
	cur = head
	curCopy = ret
	for cur != nil {
		next := cur.next.next
		curCopy = cur.next
		cur.next = next
		if next != nil {
			curCopy.next = next.next
		} else {
			curCopy.next = nil
		}
	}
	return ret
}

// 返回有环链表的一个入环节点
// 1. 遍历，把地址放进set里 发现有重复的节点时候为第一个入环节点
// 2. F指针一次走两步 S指针一次走一步，FS相遇的时候，F回开头 一次走一步 S继续一次走一步，第二次相遇即为入环点
//
func FindLoopNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	f := head.Next.Next
	s := head.Next
	for f != s {
		if s == nil || f == nil || f.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next
	}
	f = head
	for f != s {
		f = f.Next
		s = s.Next
	}
	return f
}

// 返回相交链表的第一个相交节点
// 1.用hashset
// 2.head1到最后一个不为空的节点end1  记录长度len1  head2同样操作
//   假设len1 > len2 p1从head1走len2 - len1步 然后两遍继续往前走 一次一步 相等时则为入口
func GetFirstXNodeWithoutLoop(head1 *ListNode, head2 *ListNode) *ListNode {
	cur1 := head1
	cur2 := head2
	n := 0
	// 注意用的是结尾节点的前一个
	for cur1.Next != nil {
		n++
		cur1 = cur1.Next
	}
	for cur2.Next != nil {
		n--
		cur2 = cur2.Next
	}
	if cur1 == cur2 {
		return nil
	}
	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}
	n = int(math.Abs(float64(n)))
	for n > 0 {
		cur1 = cur1.Next
		n--
	}
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

// 如果有环的
// 1.有环无交集
// 2.入环节点相同 把loop1和loop2作为终点即可 跟上面的一样
// 3.入环节点不同 loop1继续往下走，如果遇到loop2 返回任意一个即可，否则返回空
func GetFirstXNodeWithLoop(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := FindLoopNode(head1)
	loop2 := FindLoopNode(head2)
	if loop1 == nil && loop2 == nil {
		return GetFirstXNodeWithoutLoop(head1, head2)
	}
	// 如果一个有环一个无环返回nil
	if loop1 == loop2 {
		cur1 := head1
		cur2 := head2
		n := 0
		for cur1.Next != loop1 {
			n++
			cur1 = cur1.Next
		}
		for cur2.Next != loop2 {
			n--
			cur2 = cur2.Next
		}
		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}
		n = int(math.Abs(float64(n)))
		for n > 0 {
			cur1 = cur1.Next
			n--
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		p := loop1.Next
		for p != loop1 {
			if p == loop2 {
				return loop1
			}
			p = p.Next
		}
		return nil
	}
}
