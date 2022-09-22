package graphs

func DFS(matrix [][]int, i, j int, visited map[*int]bool) {
	visited[&matrix[i][j]] = true
	for _, neighbor := range getNeighbors(matrix, i, j) {
		u, v := neighbor[0], neighbor[1]
		if !visited[&matrix[u][v]] {
			DFS(matrix, u, v, visited)
		}
	}
}
