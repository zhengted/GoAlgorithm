package Zuo

// trieTree 前缀树
type Node struct {
	pass  int
	end   int
	nexts [26]*Node
}

func NewNode() *Node {
	return &Node{
		pass:  0,
		end:   0,
		nexts: [26]*Node{nil},
	}
}

type TrieTree struct {
	root *Node
}

func NewTrieTree() TrieTree {
	return TrieTree{
		root: NewNode(),
	}
}

func (t *TrieTree) insert(word string) {
	if word == "" {
		return
	}
	str := []byte(word)
	node := t.root
	node.pass++
	path := 0
	for i := 0; i < len(str); i++ {
		path = int(str[i] - 'a')
		if node.nexts[path] == nil {
			node.nexts[path] = NewNode()
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

func (t *TrieTree) search(word string) int {
	if len(word) == 0 {
		return 0
	}
	str := []byte(word)
	node := t.root
	for i := 0; i < len(str); i++ {
		index := int(str[i] - 'a')
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

func (t *TrieTree) prefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}
	str := []byte(pre)
	node := t.root
	index := 0
	for i := 0; i < len(str); i++ {
		index = int(str[i] - 'a')
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.pass
}

func (t *TrieTree) delete(word string) {
	if t.search(word) != 0 {
		// 沿途p-- 最后一个节点e--
		str := []byte(word)
		index := 0
		node := t.root
		node.pass--
		for i := 0; i < len(str); i++ {
			index = int(str[i] - 'a')
			if node.nexts[index].pass-1 == 0 {
				node.nexts[index].pass--
				node.nexts[index] = nil
				return
			}
			node = node.nexts[index]
		}
		node.end--
	}
}
