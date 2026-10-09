func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	squares := [9][9]bool{}

	for r:=0; r < 9; r++{
		for c:=0; c < 9; c++{
			val := board[r][c]
			if val == '.'{
				continue
			}
			//Map byte 1 - 9 to 0 - 8
			digitIdx := val - '1'
			squareIdx := (r/3)*3 + c/3
			if rows[r][digitIdx] || cols[digitIdx][c] || squares[squareIdx][digitIdx]{
				return false
			}
			rows[r][digitIdx] = true
			cols[digitIdx][c] = true
			squares[squareIdx][digitIdx] = true
		}
	}
	return true
}
