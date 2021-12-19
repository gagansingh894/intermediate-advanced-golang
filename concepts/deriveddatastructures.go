package concepts

import "fmt"

type Node struct {
	value int
	next *Node
	previous *Node
}

type List struct {
	head *Node
	length int
}

type Stack struct {
	elements []interface{}
}

type Queue struct {
	elements []interface{}
}

type BoundedQueue struct {
	size int
	elements chan interface{}
}

func (l *List) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Previous() *Node {
	return n.previous
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

func (s *Stack) Push(value interface{}) *Stack {
	s.elements = append(s.elements, value)
	return s
}

func (s *Stack) Pop() interface{} {
	length := len(s.elements)
	if length == 0 {
		fmt.Println("Stack is empty!")
		return nil
	}
	var element interface{}
	element, s.elements = s.elements[len(s.elements)-1], s.elements[:len(s.elements)-1]
	return element
}

func (q *Queue) Enqueue(value interface{}) *Queue {
	q.elements = append(q.elements, value)
	return q
}

func (q *Queue) Dequeue() interface{} {
	length := len(q.elements)
	if length == 0 {
		fmt.Println("Stack is empty!")
		return nil
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (qb *BoundedQueue) Enqueue(value interface{}) *BoundedQueue {
	qb.elements <- value
	return qb
}

func (qb *BoundedQueue) Dequeue() interface{} {
	return <-qb.elements
}

func newBoundedQueue(size int) *BoundedQueue{
	return &BoundedQueue{
		size:     size,
		elements: make(chan interface{}, size),
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
	s := Stack{}
	s.Push(5)
	s.Push(3)
	s.Push(2)
	fmt.Println(s.elements)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Pop()
	fmt.Println(s.elements)
}

func QueueExample(){
	q := Queue{}
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.elements)
	q.Dequeue()
	fmt.Println(q.elements)
	q.Dequeue()
	fmt.Println(q.elements)

}

func BoundedQueueExample(){
	qb := newBoundedQueue(4)
	qb.Enqueue(5)
	qb.Enqueue(5)
	fmt.Println(qb.elements)
	fmt.Println(qb.Dequeue())
	fmt.Println(qb.Dequeue())

}