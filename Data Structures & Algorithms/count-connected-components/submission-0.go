func countComponents(n int, edges [][]int) int {
    adjList := make([][]int,n)
	visited := make([]bool,n)
	for _,edge := range edges{
		u,v := edge[0],edge[1]
		adjList[u] = append(adjList[u],v)
		adjList[v] = append(adjList[v],u)
	}

	var dfs func(int)
	dfs = func(node int){
		for _,neighbor := range adjList[node]{
			if !visited[neighbor]{
				visited[neighbor] = true
				dfs(neighbor)
			}
		}
	}
	count := 0 
	for i := 0;i < n;i++{
		if !visited[i]{
			visited[i] = true
			dfs(i)
			count++
		}
	}
	return count
}
