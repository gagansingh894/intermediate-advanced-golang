package concepts

import "fmt"

type List struct {
	head *Node
	length int
}

type Node struct {
	value int
	next *Node
	previous *Node
}


func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Insert(value int)  {
	newNode := &Node{
		value: value,
		next:  nil,
		previous: nil,
	}
	if l.First() == nil {
		l.head = newNode
		l.length++
	} else {
		currentNode := l.First()
		for {
			if currentNode.next == nil {
				currentNode.next = newNode
				currentNode.next.previous = currentNode
				l.length++
				break
			}
			currentNode = currentNode.Next()
		}
	}
}


func (l *List) Traverse() {
	currentNode := l.First()
	for {
		if currentNode == nil {
			break
		}
		fmt.Println("Value: ", currentNode.value)
		fmt.Println("Next Node: ", currentNode.Next())
		fmt.Println("Previous Node: ", currentNode.Previous())
		currentNode = currentNode.Next()
	}
}

func ListExample() {
	l := List{
		head:   nil,
		length: 0,
	}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Traverse()
}