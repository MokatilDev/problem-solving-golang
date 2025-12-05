package implement_stack_using_queues

type MyStack struct {
	items []int
	size  int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
	this.size++
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	element := this.items[this.size]
	this.items = this.items[:this.size-1]
	this.size--
	return element
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.items[0]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
