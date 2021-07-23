package main

import (
	"fmt"
)

type Name interface {
	name()
}

type name1 struct {
	Name
}

//func(name1) name()  {
//	fmt.Println("kkkkooooo")
//}

type kk struct {
}

func (kk) name() {
	fmt.Println("kkkk")
}
func main() {
	//	context.WithCancel()
	//context.WithValue()
	//	ee  := name1{kk{}}
	//d.name.name()
	//fmt.Println(ll)
	//ch := make(chan int)
	//close(ch)
	//	for i := 0 ; i < 6 ;i++{
	//		v , f :=<-ch
	//		fmt.Println(v , f)
	//	}
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
