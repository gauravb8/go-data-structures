package main

type Heap struct {
	arr []int
}

func (h *Heap) swap(i, j int) {
	// fmt.Println("Swapping", h.arr[j], "and", h.arr[i])
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap) Heapify(i int) {
	l, r := left(i), right(i)
	size := len(h.arr)

	var largest int
	// fmt.Println("Node:", h.arr[i])
	if l < size {
		// fmt.Println("Left:", h.arr[l])
	}
	if r < size {
		// fmt.Println("Right:", h.arr[r])
	}
	if l < size && h.arr[l] > h.arr[i] {
		largest = l
	} else {
		largest = i
	}
	if r < size && h.arr[r] > h.arr[largest] {
		largest = r
	}

	if largest != i {
		h.swap(i, largest)
		h.Heapify(largest)
	}
}

func buildMaxHeap(a []int) *Heap {
	h := &Heap{
		arr: a,
	}

	n := len(a)

	for i := n / 2; i >= 0; i-- {
		h.Heapify(i)
	}

	return h
}

func (h *Heap) insert(elem int) {
	h.arr = append(h.arr, elem)
	for i := len(h.arr) - 1; i > 0 && h.arr[(i-1)/2] < h.arr[i]; i = (i - 1) / 2 {
		// fmt.Println("Swapping", h.arr[i], "and", h.arr[(i-1)/2])
		h.swap(i, (i-1)/2)
	}
}

func (h *Heap) extractMax() int {
	max := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.Heapify(0)
	return max
}
