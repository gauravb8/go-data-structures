package stack

type Stack struct {
	arr []int
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) push(val int) {
	s.arr = append(s.arr, val)
}

func (s *Stack) pop() int {
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem
}

func (s *Stack) peek() int {
	return s.arr[len(s.arr)-1]
}
