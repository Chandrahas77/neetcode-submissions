func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	q := make([][2]int,0)
	freshCount := 0
	time := 0
	for r := range m{
		for c := range n{
			if grid[r][c] == 1{
				freshCount++
			} 
			if grid[r][c] == 2{
				q = append(q,[2]int{r,c})
			}
		}
	}
	if freshCount == 0{
		return 0
	}
	directions := [][2]int{{-1,0},{1,0},{0,-1},{0,1},}
	for len(q) > 0 && freshCount > 0{
		qLen := len(q)
		for range qLen{
			curr := q[0]
			q = q[1:]
			r,c := curr[0],curr[1]
			for _,dir := range directions{
				nr := r+dir[0]
				nc := c+dir[1]
				if nr >= 0 && nc >= 0 && nr < m && nc < n && grid[nr][nc] == 1{
					grid[nr][nc] = 2
					q = append(q,[2]int{nr,nc})
					freshCount--
				}
			}
		}
		time++
	}
	if freshCount == 0{
		return time
	}
	return -1
}
