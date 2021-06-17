package Zuo

import (
	. "GoAlgorithm/Common"
	"container/heap"
	"fmt"
	"math"
	"math/rand"
)

// GenerateGraph: 给定一个二维矩阵 [0] 表示边权重 [1] 表示from的节点value [2]表示to节点的value
func GenerateGraph(matrix [][]int) Graph {
	g := NewGraph()
	for _, edge := range matrix {
		weight := edge[0]
		from := edge[1]
		to := edge[2]
		if _, ok := g.Nodes[from]; !ok {
			g.Nodes[from] = NewGraphNode(from)
		}
		if _, ok := g.Nodes[to]; !ok {
			g.Nodes[to] = NewGraphNode(to)
		}
		nFrom := g.Nodes[from]
		nTo := g.Nodes[to]
		newEdge := NewGraphEdge(weight, nFrom, nTo)
		g.Edges = append(g.Edges, newEdge)
		nFrom.Out++
		nTo.In++
		nFrom.Edges = append(nFrom.Edges, newEdge)
		nFrom.Nexts = append(nFrom.Nexts, nTo)
	}
	return g
}

// bfs: 图的广度优先遍历
func bfs(node *GraphNode) {
	q := []*GraphNode{}
	s := []*GraphNode{}
	q = append(q, node)
	s = append(s, node)
	for len(q) > 0 {
		cur := q[0]
		if len(q) > 1 {
			q = q[1:]
		} else {
			q = []*GraphNode{}
		}
		fmt.Println(cur.Value)
		for _, next := range cur.Nexts {
			if !isNodeInSet(s, next) {
				q = append(q, next)
				s = append(s, next)
			}
		}
	}
}

// dfs:图的深度优先遍历
func dfs(node *GraphNode) {
	if node == nil {
		return
	}
	stack := []*GraphNode{}
	set := []*GraphNode{}
	stack = append(stack, node)
	// 登记就打印
	set = append(set, node)
	fmt.Println(node.Value)
	for len(stack) > 0 {
		cur := stack[0]
		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = []*GraphNode{}
		}
		for _, next := range cur.Nexts {
			if !isNodeInSet(set, next) {
				stack = append(stack, cur)
				stack = append(stack, next)
				// 登记
				set = append(set, next)
				fmt.Println(next.Value)
				break
			}
		}
	}
}

func isNodeInSet(set []*GraphNode, node *GraphNode) bool {
	for i := 0; i < len(set); i++ {
		if set[i] == node {
			return true
		}
	}
	return false
}

// topo:图的拓扑排序，有向无环图
func topo(graph Graph) []*GraphNode {
	m := make(map[*GraphNode]int) // 每个节点的入度集合
	q := []*GraphNode{}
	for _, graphNode := range graph.Nodes {
		m[graphNode] = graphNode.In
		if graphNode.In == 0 {
			q = append(q, graphNode)
		}
	}

	res := []*GraphNode{}
	for len(q) > 0 {
		cur := q[0]
		if len(q) > 1 {
			q = q[1:]
		} else {
			q = []*GraphNode{}
		}
		res = append(res, cur)
		for _, next := range cur.Nexts {
			next.In--
			if next.In == 0 {
				q = append(q, next)
			}
		}
	}
	return res
}

//无向图 最小生成树
// Kruskal: 先排序(小顶堆)，按序取边，计算并查集

// 并查集部分 用来存节点的
type GraphNodeUnionSet struct {
	parent  map[*GraphNode]*GraphNode // 父节点集合
	sizeMap map[*GraphNode]int        // 每个父节点所在的集合的大小
}

func (s GraphNodeUnionSet) IsSameSet(a, b *GraphNode) bool {
	if s.findFather(a) == s.findFather(b) {
		return true
	}
	return false
}

func (s *GraphNodeUnionSet) Union(a, b *GraphNode) {
	ap := s.findFather(a)
	bp := s.findFather(b)
	aSize := s.sizeMap[ap]
	bSize := s.sizeMap[bp]
	if aSize > bSize {
		s.parent[bp] = ap
		s.sizeMap[ap] = aSize + bSize
		delete(s.sizeMap, bp)
	} else {
		s.parent[ap] = bp
		s.sizeMap[bp] = aSize + bSize
		delete(s.sizeMap, ap)
	}

}

func (s *GraphNodeUnionSet) findFather(node *GraphNode) *GraphNode {
	temp := []*GraphNode{}
	for node != s.parent[node] {
		temp = append(temp, node)
		node = s.parent[node]
	}
	for _, graphNode := range temp {
		s.parent[graphNode] = node
	}
	return node

}

type GraphEdgeExSt struct {
	*GraphEdge
	Index int
}

// PQ部分用来pick最小边的
type GraphEdgePQ []*GraphEdgeExSt

func (pq GraphEdgePQ) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq GraphEdgePQ) Len() int {
	return len(pq)
}

func (pq GraphEdgePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = j
	pq[j].Index = i
}

func (pq *GraphEdgePQ) Push(edge interface{}) {
	edgeEx := edge.(GraphEdgeExSt)
	n := len(*pq)
	edgeEx.Index = n
	*pq = append(*pq, &edgeEx)
}

func (pq *GraphEdgePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.Index = -1 // 外部无用字段
	old[n-1] = nil // 避免内存泄露
	old = old[:n-1]
	*pq = old
	return ret
}

func (pq *GraphEdgePQ) Update(item *GraphEdgeExSt, index int) {
	(*pq)[index].Weight = item.Weight
	heap.Fix(pq, index)
}

func Kruskal(graph *Graph) int {
	res := 0
	if graph == nil {
		return res
	}
	nodeUnionSet := GraphNodeUnionSet{
		parent:  make(map[*GraphNode]*GraphNode),
		sizeMap: make(map[*GraphNode]int),
	}
	for _, graphNode := range graph.Nodes {
		nodeUnionSet.parent[graphNode] = graphNode
		nodeUnionSet.sizeMap[graphNode] = 1
	}

	edgePq := &GraphEdgePQ{}
	heap.Init(edgePq)

	for _, edge := range graph.Edges {
		edgeEx := &GraphEdgeExSt{edge, -1}
		heap.Push(edgePq, edgeEx)
	}

	for edgePq.Len() > 0 {
		curEdgeEx := edgePq.Pop().(GraphEdgeExSt)
		nodeFrom := curEdgeEx.From
		nodeTo := curEdgeEx.To
		if !nodeUnionSet.IsSameSet(nodeFrom, nodeTo) {
			nodeUnionSet.Union(nodeFrom, nodeTo)
			res += curEdgeEx.Weight
		}
	}

	return res
}

// Prim: 出发点-->解锁周围的周边-->选个最小的边，解锁他的出点-->解锁这个出点周围的边。。。。（解锁成功后拿个点集存储）
func Prim(graph *Graph) int {
	res := 0
	if graph == nil {
		return res
	}
	edgePq := &GraphEdgePQ{}
	heap.Init(edgePq)
	//selectedEdge := make(map[*GraphEdge]int)
	selectedNode := make(map[*GraphNode]int)
	for _, node := range graph.Nodes {
		if _, ok := selectedNode[node]; ok {
			continue
		}
		selectedNode[node] = 1
		for _, edge := range node.Edges {
			edgeEx := &GraphEdgeExSt{edge, -1}
			heap.Push(edgePq, edgeEx)
		}
		// 这个循环一开始以为不会跳出，其实
		for edgePq.Len() > 0 {
			curEdgeEx := heap.Pop(edgePq).(GraphEdgeExSt)
			to := curEdgeEx.To
			if _, ok := selectedNode[to]; !ok {
				selectedNode[node] = 1
				res += curEdgeEx.Weight // 这里没考虑无向图因为是双向的 这条边不会被加两次，因为上面判断了这个节点是不是已经被PICK过了

				for _, edge := range node.Edges {
					edgeEx := &GraphEdgeExSt{edge, -1}
					heap.Push(edgePq, edgeEx)
				}
			}
		}
		break
	}
	return res
}

// 有向图 最短路径
// Dijkstra：解决起始点到其他点的最短距离问题
//	优化小顶堆Update问题
func Dijkstra(node *GraphNode) (ret map[*GraphNode]int) {
	ret = map[*GraphNode]int{}
	ret[node] = 0
	lockedNode := map[*GraphNode]int{}
	minNode := getMinDistanceWithUnselected(ret, lockedNode)
	for nil != minNode {
		for _, edge := range minNode.Edges {
			if _, ok := ret[edge.To]; !ok {
				ret[edge.To] = ret[minNode] + edge.Weight
			} else {
				ret[edge.To] = MinInt(ret[edge.To], ret[minNode]+edge.Weight)
			}
		}
		lockedNode[minNode] = 1
		minNode = getMinDistanceWithUnselected(ret, lockedNode)
	}
	return
}

// 找出未锁定节点中distance最小的
func getMinDistanceWithUnselected(distanceMap, lockedNode map[*GraphNode]int) (retNode *GraphNode) {
	retMin := math.MaxInt32
	for graphNode, nDis := range distanceMap {
		if _, ok := lockedNode[graphNode]; !ok {
			if retMin > nDis {
				retMin = nDis
				retNode = graphNode
			}
		}
	}
	return retNode
}

func QuickSort0615(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	newL, newR := partition(arr, l, r)
	quickSort(arr, l, newL-1)
	quickSort(arr, newR+1, r)
}

func partition(arr []int, l, r int) (int, int) {
	randNum := l + rand.Intn(r-l)
	arr[randNum], arr[r] = arr[r], arr[randNum]
	less := l - 1
	more := r
	index := l
	for index < more {
		if index < arr[r] {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			index++
			less++
		} else if index > arr[r] {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		} else {
			index++
		}
	}
	arr[more], arr[r] = arr[r], arr[more]
	return less + 1, more
}
