func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
	numIsland := 0

	var dfs func(r,c int)
	dfs = func(r,c int){
		if r < 0 || c < 0 || r >= m || c >= n || grid[r][c] == '0'{
			return
		}
		//set the visited island to 0 so we dont revisit them again
		grid[r][c] = '0'
		dfs(r-1,c)
		dfs(r+1,c)
		dfs(r,c-1)
		dfs(r,c+1)
	}
	for r := 0; r < m; r++{
		for c := 0; c < n; c++{
			if grid[r][c] == '1'{
				//recursively set the neighboring islands to 0 so we don't visit them again
				dfs(r,c)
				numIsland++
			}
		}
	}
	return numIsland
}
