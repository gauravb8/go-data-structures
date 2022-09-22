package graphs

import "github.com/gauravb8/go-data-structures/queue"

func DFS(matrix [][]int, i, j int, visited map[*int]bool) {
	// Perform Depth-first search recursively starting at cell (i,j)
	visited[&matrix[i][j]] = true
	for _, neighbor := range getNeighbors(matrix, i, j) {
		u, v := neighbor[0], neighbor[1]
		if !visited[&matrix[u][v]] {
			DFS(matrix, u, v, visited)
		}
	}
}

func BFS(matrix [][]int, i, j int) {
	// Perform Breadth-first search on matrix starting with cell (i,j)
	q := &queue.Queue{}
	visited := make(map[*int]bool)

	q.Push([]int{i, j})
	visited[&matrix[i][j]] = true

	for !q.IsEmpty() {
		n := q.Pop()
		u, v := n[0], n[1]

		for _, neighbor := range getNeighbors(matrix, u, v) {
			x, y := neighbor[0], neighbor[1]
			if !visited[&matrix[x][y]] {
				q.Push([]int{x, y})
				visited[&matrix[x][y]] = true
			}
		}
	}
}
