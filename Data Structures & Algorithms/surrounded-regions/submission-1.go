func solve(board [][]byte) {
    m, n := len(board), len(board[0])

	var capture func(r,c int)
	capture = func(r,c int){
		if r < 0 || c < 0 || r >= m || c >= n || board[r][c] != 'O'{
			return
		}
		board[r][c] = 'T'
		capture(r-1,c)
		capture(r+1,c)
		capture(r,c-1)
		capture(r,c+1)
	}
	//capture un-surrounded regions (DFS)
	for r := range m{
		for c := range n{
			if board[r][c] == 'O' && (r == 0 || r == m-1 || c == 0 || c == n-1){
				capture(r,c)
			}
		}
	}
	//Capture surrounded regions
	for r := range m{
		for c := range n{
			if board[r][c] == 'O'{
				board[r][c] = 'X'
			}
		}
	}
	//Uncapture unsurrounded regions
	for r := range m{
		for c := range n{
			if board[r][c] == 'T'{
				board[r][c] = 'O'
			}
		}
	}
}
