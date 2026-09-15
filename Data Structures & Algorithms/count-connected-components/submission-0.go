func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	components := 0

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, neighbor := range graph[node]{
			if !visited[neighbor]{
				dfs(neighbor)
			}
		}
	}

	for node := 0; node < n; node++{
		if !visited[node] {
			components++
			dfs(node)
		}
	}

	return components
}
