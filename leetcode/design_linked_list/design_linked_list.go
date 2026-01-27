package design_linkedlist

type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.Head == nil || index < 0 {
		return -1
	}

	current := this.Head
	for i := 0; i < index; i++ {
		if current.Next == nil {
			return -1
		}
		current = current.Next
	}

	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val, Next: this.Head}
	this.Head = newNode

	if this.Tail == nil {
		this.Tail = newNode
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}
	if this.Tail == nil {
		this.Tail = newNode
		this.Head = newNode
		return
	}

	this.Tail.Next = newNode
	this.Tail = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}

	if current == nil {
		return
	}

	newNode := &ListNode{Val: val, Next: current.Next}
	current.Next = newNode
	if newNode.Next == nil {
		this.Tail = newNode
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.Head == nil {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Head == nil {
			this.Tail = nil
		}
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		if current.Next == nil {
			return
		}
		current = current.Next
	}

	if current.Next == nil {
		return
	}

	if current.Next == this.Tail {
		this.Tail = current
	}

	current.Next = current.Next.Next
}
