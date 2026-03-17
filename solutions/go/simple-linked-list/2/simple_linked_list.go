package linkedlist

import "errors"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head   *Node
	Length int
}

func New(array []int) *List {
	var list List
	for _, data := range array {
		list.Push(data)
	}
	return &list
}

func (list *List) Size() int {
	return list.Length
}

func (list *List) Push(data int) {
	newNode := Node{data, nil}
	if list.Head == nil {
		list.Head = &newNode
	} else {
		node := list.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = &newNode
	}
	list.Length++
}

func (list *List) Pop() (int, error) {
	var data int
	switch list.Length {
	case 0:
		return 0, errors.New("can not pop from empty list")
	case 1:
		data = list.Head.Data
		list.Head = nil
	default:
		node := list.Head
		for node.Next.Next != nil {
			node = node.Next
		}
		data = node.Next.Data
		node.Next = nil
	}
	list.Length--
	return data, nil
}

func (list *List) Array() []int {
	arr := make([]int, list.Length)
	for i, node := 0, list.Head; node != nil; node = node.Next {
		arr[i] = node.Data
		i++
	}
	return arr
}

func (list *List) Reverse() *List {
	arr, length := list.Array(), list.Length
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
	}
	return New(arr)
}
