package stack

type Stack struct {
	arr []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() int {
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem
}

func (s *Stack) Peek() int {
	return s.arr[len(s.arr)-1]
}
