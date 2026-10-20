func pacificAtlantic(heights [][]int) [][]int {
    m,n := len(heights),len(heights[0])
	pacific := make([][]bool,m)
	atlantic := make([][]bool,m)
	for i := range m{
		pacific[i] = make([]bool,n)
		atlantic[i] = make([]bool,n)
	}
	var dfs func(r,c int,ocean [][]bool,prevHeight int)
	dfs = func(r,c int,ocean [][]bool,prevHeight int){
		if r < 0 || c < 0 || r >= m || c >= n || ocean[r][c] || heights[r][c] < prevHeight{
			return
		}
		ocean[r][c] = true
		dfs(r-1,c,ocean,heights[r][c])
		dfs(r+1,c,ocean,heights[r][c])
		dfs(r,c-1,ocean,heights[r][c])
		dfs(r,c+1,ocean,heights[r][c])
	}
	//vertical ends
	for r := range m{
		dfs(r,0,pacific,heights[r][0])
		dfs(r,n-1,atlantic,heights[r][n-1])
	}
	//horizontal ends
	for c := range n{
		dfs(0,c,pacific,heights[0][c])
		dfs(m-1,c,atlantic,heights[m-1][c])
	}
	res := make([][]int,0)
	for r := range m{
		for c := range n{
			if pacific[r][c] && atlantic[r][c]{
				res = append(res,[]int{r,c})
			}
		}
	}
return res
}
