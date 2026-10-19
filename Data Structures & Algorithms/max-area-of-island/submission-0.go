func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
	n := len(grid[0])
	maxArea := 0

	var dfs func(r,c int) int
	dfs = func(r,c int) int{
		if r < 0 || c < 0 || r >= m || c >= n || grid[r][c] == 0{
			return 0
		}
		grid[r][c] = 0
		return 1 + dfs(r-1,c) + dfs(r+1,c) + dfs(r,c-1) + dfs(r,c+1)
	}
	for r := 0; r < m; r++{
		for c := 0; c < n; c++{
			if grid[r][c] == 1{
				area := dfs(r,c)
				if area > maxArea{
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
