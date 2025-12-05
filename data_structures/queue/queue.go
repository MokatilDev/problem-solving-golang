package queue

import "errors"

type Queue struct {
	items []int
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) isEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue(data int) (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Underflow")
	}

	element := q.items[0]
	if q.Len() == 1 {
		q.items = nil
		return element, nil
	}
	q.items = q.items[1:]
	return element, nil
}

func (q *Queue) Peek() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Empty queue")
	}

	return q.items[0], nil
}
