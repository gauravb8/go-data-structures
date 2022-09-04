package queue

import "errors"

// Double ended queue
type Deque struct {
	arr []int
}

func (q *Deque) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Deque) Size() int {
	return len(q.arr)
}

func (q *Deque) PopFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is Empty")
	}

	elem := q.arr[0]
	q.arr = q.arr[1:]
	return elem, nil
}

func (q *Deque) PopRear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is Empty")
	}

	last := len(q.arr) - 1
	elem := q.arr[last]
	q.arr = q.arr[:last]
	return elem, nil
}

func (q *Deque) PushRear(val int) {
	q.arr = append(q.arr, val)
}
