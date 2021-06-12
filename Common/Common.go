package Common

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Pre  *DoubleNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Swap:只有内存不相同的两个变量才可以使用
// 别炫技，只当拓展
func Swap(i, j *int) {
	*i = *i ^ *j
	*j = *i ^ *j
	*i = *i ^ *j
}
