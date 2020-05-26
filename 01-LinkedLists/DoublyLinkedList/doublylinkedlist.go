package main

import "fmt"

// Node struct that keeps track of next node and value of node
type Node struct {
	Next  *Node
	Prev  *Node
	Value string
}

// LinkedList datastructure to manage head, define functions
type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) append(val string) error {
	nodeToAdd := &Node{Next: nil, Prev: nil, Value: val}
	if l.Head == nil {
		l.Head = nodeToAdd
		l.Tail = nodeToAdd
		return nil
	}
	nodeToAdd.Prev = l.Tail
	l.Tail.Next = nodeToAdd
	l.Tail = nodeToAdd
	return nil
}

func (l *LinkedList) insert(index int, value string) error {
	counter := 0
	nodeToAdd := &Node{Next: nil, Prev: nil, Value: value}
	if index == 0 {
		nodeToAdd.Next = l.Head
		if l.Head != nil {
			l.Head.Prev = nodeToAdd
		}
		l.Head = nodeToAdd
		return nil
	}
	curr := l.Head
	for curr != nil {
		if index == counter {
			prev := curr.Prev
			nodeToAdd.Prev = prev
			prev.Next = nodeToAdd
			nodeToAdd.Next = curr
			curr.Prev = nodeToAdd
			return nil
		}
		curr = curr.Next
		counter++
	}
	return nil
}

func (l *LinkedList) delete(index int) error {
	if l.Head == nil {
		return nil
	}
	if index == 0 {
		oldHead := l.Head
		l.Head = oldHead.Next
		oldHead.Next = nil
		return nil
	}
	counter := 0
	curr := l.Head
	for curr != nil {
		if counter == index {
			curr.Prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = curr.Prev
			} else {
				curr.Next.Prev = curr.Prev
			}
			return nil
		}
		curr = curr.Next
		counter++
	}
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

func (l *LinkedList) printReversed() {
	str := ""
	curr := l.Tail
	for curr != nil {
		str += (*curr).Value + " -> "
		curr = curr.Prev
	}
	str += "nil"
	fmt.Println(str)
}

func main() {
	var linkedList LinkedList
	linkedList.append("a")
	linkedList.append("c")
	linkedList.append("d")
	linkedList.append("e")
	linkedList.append("f")
	linkedList.insert(1, "b")
	linkedList.insert(7, "g")
	linkedList.print()
	fmt.Println("Printing reversed: ")
	linkedList.printReversed()
	fmt.Println("Deleting: ")
	linkedList.delete(0)
	linkedList.delete(4)
	linkedList.print()
}
