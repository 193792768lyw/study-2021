package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {

	newHead := &Node{}
	old := head
	pre := newHead
	m := make(map[*Node]int)
	mNew := make(map[int]*Node)
	index := 1
	for old != nil {
		node := &Node{
			Val:    old.Val,
			Next:   nil,
			Random: nil,
		}

		m[old] = index
		mNew[index] = node
		index++

		pre.Next = node
		pre = pre.Next
		old = old.Next
	}

	old = head
	pre = newHead.Next
	for old != nil {
		if old.Random != nil {
			pre.Random = mNew[m[old.Random]]
		}
		pre = pre.Next
		old = old.Next
	}

	return newHead.Next

}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}
