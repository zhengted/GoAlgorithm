package Zuo

// 并查集
//union：小集合的头-->大集合的头
//isSet：向上找头，头相同return true

type node struct {
	val interface{}
}

type UnionSet interface {
	findFather(cur *node) *node
	isSameSet(a, b *node) bool
	union(a, b *node)
	getNode(val interface{}) *node
	containsKey(val interface{}) bool
}

type UnionSetSt struct {
	nodes   map[interface{}]*node
	parents map[*node]*node
	sizeMap map[*node]int
}

func NewUnionSet(list []interface{}) *UnionSetSt {
	ret := &UnionSetSt{
		nodes:   make(map[interface{}]*node),
		parents: make(map[*node]*node),
		sizeMap: make(map[*node]int),
	}
	for _, val := range list {
		n := &node{val: val}
		ret.nodes[val] = n
		ret.sizeMap[n] = 1
		ret.parents[n] = n
	}
	return ret
}

func (s UnionSetSt) getNode(val interface{}) *node {
	return s.nodes[val]
}

func (s UnionSetSt) findFather(n *node) *node {
	// 优化 寻找祖宗节点的过程，将沿途节点变为最终祖宗节点的直接子节点
	temp := []*node{}
	for n != s.parents[n] {
		temp = append(temp, n)
		n = s.parents[n]
	}
	for _, node := range temp {
		s.parents[node] = n
	}
	return n
}

func (s UnionSetSt) containsKey(val interface{}) bool {
	_, ok := s.nodes[val]
	return ok
}

func (s UnionSetSt) isSameSet(a, b interface{}) bool {
	if !s.containsKey(a) || !s.containsKey(b) {
		return false
	}
	return s.findFather(s.getNode(a)) == s.findFather(s.getNode(b))
}

func (s *UnionSetSt) union(a, b *node) {
	if !s.containsKey(a.val) || !s.containsKey(b.val) {
		return
	}
	ahead := s.findFather(a)
	bhead := s.findFather(b)
	if ahead != bhead {
		aSetSize := s.sizeMap[ahead]
		bSetSize := s.sizeMap[bhead]
		if aSetSize > bSetSize {
			s.parents[bhead] = ahead
			s.sizeMap[ahead] = aSetSize + bSetSize
			delete(s.sizeMap, bhead)
		} else {
			s.parents[ahead] = bhead
			s.sizeMap[bhead] = aSetSize + bSetSize
			delete(s.sizeMap, bhead)
		}
	}
}

// 每个用户有三个字段，两个用户只要有一个字段相同可以认为是同一个 问 给出一个用户数组
// 给出一个用户数组，判断实际用户个数
type User struct {
	a int
	b int
	c int
}
