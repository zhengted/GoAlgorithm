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

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type GraphNode struct {
	Value int          // 编号
	In    int          // 入度
	Out   int          // 出度
	Nexts []*GraphNode // 直接邻居，表示以该点出发的
	Edges []*GraphEdge // 直接边，同样是以该结点出发
}

type GraphEdge struct {
	Weight int
	From   *GraphNode
	To     *GraphNode
}

type Graph struct {
	Nodes map[int]*GraphNode
	Edges []*GraphEdge
}

func NewGraphNode(index int) *GraphNode {
	return &GraphNode{
		Value: index,
		In:    0,
		Out:   0,
		Nexts: []*GraphNode{},
		Edges: []*GraphEdge{},
	}
}

func NewGraphEdge(weight int, from, to *GraphNode) *GraphEdge {
	return &GraphEdge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

func NewGraph() Graph {
	return Graph{
		Nodes: map[int]*GraphNode{},
		Edges: []*GraphEdge{},
	}
}
