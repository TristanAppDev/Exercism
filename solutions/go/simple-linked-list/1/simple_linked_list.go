package linkedlist

import (
	"errors"
)

// Node contans its value and pointer to next node
type Node struct {
	Value int
	Next  *Node
}

// List contains a head node and a slice of nodes and a size
type List struct {
	Head  *Node
	Nodes []*Node
	size  int
}

func New(elements []int) *List {
	list := List{
		Head:  &Node{},
		Nodes: make([]*Node, len(elements)),
		size:  0,
	}

	for i, v := range elements {
		list.Nodes[i] = &Node{
			Value: v,
			Next:  nil,
		}
		if i == 0 {
			list.Head = list.Nodes[i]
		} else {
			list.Nodes[i-1].Next = list.Nodes[i]
		}
		list.size++
	}
	return &list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node := Node{
		Value: element,
		Next:  nil,
	}
	if l.size > 0 {
		l.Nodes[l.size-1].Next = &node
	}

	l.Nodes = append(l.Nodes, &node)
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("pop from empty list not allowed")
	}

	popped := l.Nodes[l.size-1]
	l.Nodes = l.Nodes[:l.size-1]
	l.size--

	if l.size > 1 {
		l.Nodes[l.size-1].Next = nil
	} else {
		l.Head = nil
	}

	return popped.Value, nil
}

func (l *List) Array() []int {
	array := make([]int, len(l.Nodes))
	for i, n := range l.Nodes {
		array[i] = n.Value
	}
	return array
}

func (l *List) Reverse() *List {
	if l.size == 0 {
		return l
	}

	newHead := l.Nodes[l.size-1]
	newHead.Next = l.Nodes[l.size-2]
	newList := List{
		Head:  l.Nodes[l.size-1],
		Nodes: make([]*Node, l.size),
		size:  l.size,
	}

	for i := 0; i < l.size; i++ {
		currentNode := l.Nodes[l.size-1-i]
		newList.Nodes[i] = &Node{currentNode.Value, nil}
		if condition := i < l.size-1; condition {
			newList.Nodes[i].Next = l.Nodes[l.size-2-i]
		}
	}

	return &newList
}
