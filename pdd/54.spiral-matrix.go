package pdd

// spiralOrder 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	visited := make(map[int]bool)
	row, col := 0, 0
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx := 0
	for i := 0; i < len(res); i++ {
		res[i] = matrix[row][col]
		visited[matrix[row][col]] = true
		row += direct[idx][0]
		col += direct[idx][1]
		if row >= m {
			row--
		}
		if col >= n {
			col--
		}
		if (row == 0 && col == n-1) || (row == m-1 && col == n-1) || (row == m-1 && col == 0) {
			idx = (idx + 1) % 4
		} else if visited[matrix[row][col]] {
			row -= direct[idx][0]
			col -= direct[idx][1]
			idx = (idx + 1) % 4
			row += direct[idx][0]
			col += direct[idx][1]
		}
	}
	return res
}
