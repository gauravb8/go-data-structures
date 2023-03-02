package graphs

import "github.com/gauravb8/go-data-structures/stack"

// Perform a topological sort on a directed graph
func TopoSort(g map[int][]int, s int, visited map[int]bool, st *stack.Stack) {
	visited[s] = true
	for _, u := range g[s] {
		if !visited[u] {
			TopoSort(g, u, visited, st)
		}
	}
	st.Push(s)
}
