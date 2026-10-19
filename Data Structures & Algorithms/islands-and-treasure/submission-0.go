func islandsAndTreasure(grid [][]int) {
    m, n := len(grid), len(grid[0])
	const INF = 2147483647
	q := make([][2]int,0)

	for r := range m{
		for c := range n{
			if grid[r][c] == 0{
				q = append(q,[2]int{r,c})
			}
		}
	}
	directions := [][2]int{
		{-1,0},
		{1,0},
		{0,-1},
		{0,1},
	}
	for len(q) > 0{
		curr := q[0]
		//pop the queue
		q = q[1:]
		r,c := curr[0],curr[1]
		for _,dir := range directions{
			nr,nc := r+dir[0],c+dir[1]
			if nr >= 0 && nc >= 0 && nr < m && nc < n && grid[nr][nc] == INF{
				grid[nr][nc] = grid[r][c] + 1
				q = append(q,[2]int{nr,nc})
			}
		}
	}
}
