package queue

type Queue struct {
	arr   []int
	front int
	rear  int
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Push(n int) {
	q.arr = append(q.arr, n)
}

func (q *Queue) Pop() int {
	i := q.arr[0]
	q.arr = q.arr[1:]
	return i
}
