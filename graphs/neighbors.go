package graphs

func getNeighbors(matrix [][]int, i, j int) [][]int {
	var neighbors [][]int
	if i > 0 {
		neighbors = append(neighbors, []int{i - 1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, []int{i, j - 1})
	}
	if i < len(matrix)-1 {
		neighbors = append(neighbors, []int{i + 1, j})
	}
	if j < len(matrix[0])-1 {
		neighbors = append(neighbors, []int{i, j + 1})
	}

	return neighbors
}
