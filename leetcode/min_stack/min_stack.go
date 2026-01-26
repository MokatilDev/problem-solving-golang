package min_stack

import (
	"math"
)

type MinStack struct {
	values [][]int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.IsEmpty() {
		this.values = append(this.values, []int{val, val})
		return
	}

	this.values = append(this.values, []int{val, int(math.Min(float64(val), float64(this.GetMin())))})
}

func (this *MinStack) IsEmpty() bool {
	return len(this.values) == 0
}

func (this *MinStack) Pop() {
	if !this.IsEmpty() {
		this.values = this.values[:len(this.values)-1]
	}
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1][0]
}

func (this *MinStack) Peak() []int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.Peak()[1]
}
