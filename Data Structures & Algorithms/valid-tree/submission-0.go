func validTree(n int, edges [][]int) bool {
    if n == 0{
		return true
	}
	adjL := make(map[int][]int,n)

	for _,edge := range edges{
		u,v := edge[0],edge[1]
		adjL[u] = append(adjL[u],v)
		adjL[v] = append(adjL[v],u)
	}
	visit := make(map[int]struct{})
	var dfs func(i,prev int) bool
	dfs = func(i,prev int) bool{
		if _,exists := visit[i]; exists{
			return false
		}
		visit[i] = struct{}{}
		for _,j := range adjL[i]{
			if j == prev{
				continue
			}
			if !dfs(j,i){
				return false
			}
		}
		return true
	}
	return dfs(0,-1) && n == len(visit)
}
