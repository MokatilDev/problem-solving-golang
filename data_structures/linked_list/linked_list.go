package linked_list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Push(val int) {
	newNode := &ListNode{Value: val}

	if l.Head == nil {
		l.Tail = newNode
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) Insert(val int, index int) {
	newNode := &ListNode{Value: val}

	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		l.Tail = newNode
		return
	}

	current := l.Head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current != nil {
		newNode.Next = current.Next
		current.Next = newNode
	}
}

func (l *LinkedList) Pop() {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	l.Tail.Next = current
	l.Tail = current
}

func (l *LinkedList) Remove(index int) {
	if l.Head == nil {
		return
	}

	if index == 0 {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *LinkedList) Print() {
	current := l.Head
	var s strings.Builder
	for current != nil {
		fmt.Fprintf(&s, "%d ->", current.Value)
	}
}
