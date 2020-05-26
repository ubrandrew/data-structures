package main

import "fmt"

// LinkedList datastructure to manage head, define functions
type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) append(val string) error {

	nodeToAdd := &Node{Next: nil, Prev: nil, Value: val}
	curr := l.Head

	if curr == nil {
		l.Head = nodeToAdd
		return nil
	}
	for (*curr).Next != nil {
		curr = (*curr).Next
	}
	(*curr).Next = nodeToAdd

	return nil
}

func (l *LinkedList) print() {
	str := ""
	curr := l.Head
	for curr != nil {
		str += (*curr).Value + " -> "
		curr = curr.Next
	}
	str += "nil"
	fmt.Println(str)
}

// Node component that keeps track of next node and value of node
type Node struct {
	Next  *Node
	Prev  *Node
	Value string
}

func main() {
	var linkedList LinkedList
	linkedList.append("a")
	linkedList.append("b")
	linkedList.append("a")
	linkedList.append("c")
	linkedList.append("a")
	linkedList.append("d")
	linkedList.print()
}
