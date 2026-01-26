package stack

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:s.Len()-1]
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.items[s.Len()-1], nil
}

func (s *Stack) Print() {
	for _, v := range s.items {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
