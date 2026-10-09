func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	top, bot := 0, rows-1
	var targetRow int

	// Binary search 1: Locate the target row
	for top <= bot {
		midRow := top + (bot-top)/2
		if target > matrix[midRow][cols-1] {
			top = midRow + 1
		} else if target < matrix[midRow][0] {
			bot = midRow - 1
		} else {
			targetRow = midRow
			break
		}
	}

	// Target lies entirely outside the matrix boundaries
	if top > bot {
		return false
	}

	// Binary search 2: Search within the identified row
	left, right := 0, cols-1
	for left <= right {
		mid := left + (right-left)/2
		val := matrix[targetRow][mid]

		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}