package main

import "fmt"

// LinkedList datastructure to manage head, define functions
type LinkedList struct {
	Head *Node
}

func (l *LinkedList) append(val string) error {

	nodeToAdd := &Node{Next: nil, Value: val}
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

func (l *LinkedList) insert(index int, val string) error {
	newNode := &Node{Next: nil, Value: val}
	counter := 0
	curr := l.Head
	var prev *Node
	// case where curr is head
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return nil
	}
	for curr != nil {
		if counter == index {
			prev.Next = newNode
			newNode.Next = curr
			return nil
		}
		prev = curr
		curr = curr.Next
		counter++
	}
	// case where curr is tail
	if index == counter+1 {
		prev.Next = newNode
	}

	return nil
}

func (l *LinkedList) delete(index int) error {
	counter := 0
	if index == 0 {
		l.Head = l.Head.Next
		return nil
	}

	curr := l.Head
	var prev *Node
	for curr != nil {
		if counter == index {
			prev.Next = curr.Next
			curr.Next = nil
		}
		prev = curr
		curr = curr.Next
		counter++
	}

	return nil

}

func (l *LinkedList) reverse() error {
	curr := l.Head
	var prev *Node
	if (*l.Head).Next == nil {
		return nil
	}
	for (*curr).Next != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev
	l.Head = curr

	return nil
}

// Node struct that keeps track of next node and value of node
type Node struct {
	Next  *Node
	Value string
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
	fmt.Println("Deleting: ")
	linkedList.delete(0)
	linkedList.delete(5)
	linkedList.print()
	fmt.Println("Reversing: ")
	linkedList.reverse()
	linkedList.print()
}
