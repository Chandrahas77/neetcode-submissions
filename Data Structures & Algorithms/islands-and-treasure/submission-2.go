func islandsAndTreasure(grid [][]int) {
    if len(grid) == 0 || len(grid[0]) == 0{
		return
	}
	m, n := len(grid),len(grid[0])
	const INF = 2147483647
	q := make([][2]int,0)
	for r := range m{
		for c := range n{
			if grid[r][c] == 0{
				q = append(q,[2]int{r,c})
			}
		}
	}
	visit := func(nr,nc int,dist int){
		if nr < 0 || nc < 0 || nr >= m || nc >=n || grid[nr][nc] != INF{
			return
		}
		grid[nr][nc] = dist
		q = append(q,[2]int{nr,nc})
	}
	for len(q) > 0{
		curr := q[0]
		q = q[1:]
		r,c := curr[0],curr[1]
		nextDist := grid[r][c] + 1
		visit(r-1,c,nextDist)	//up
		visit(r+1,c,nextDist)	//down
		visit(r,c-1,nextDist)	//left
		visit(r,c+1,nextDist)	//right
	}
}
