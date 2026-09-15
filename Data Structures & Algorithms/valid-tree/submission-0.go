func validTree(n int, edges [][]int) bool {
    if len(edges) > n-1 {
		return false
	}

	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visit := make(map[int]bool)
	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		if visit[node] {
			return false
		}
		visit[node] = true
		for _, nei := range adj[node] {
			if nei == parent {
				continue
			}
			if !dfs(nei, node) {
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && len(visit) == n
}
