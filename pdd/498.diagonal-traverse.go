package pdd

// findDiagonalOrder 498. 对角线遍历
func findDiagonalOrder(mat [][]int) []int {
	if mat == nil || len(mat) == 0 {
		return nil
	}

	m, n := len(mat), len(mat[0])
	row, col := 0, 0
	res := make([]int, m*n)
	for i := 0; i < len(res); i++ {
		res[i] = mat[row][col]
		// row+col 即为遍历的层数，偶数是向上遍历，奇数是向下遍历
		if ((row + col) % 2) == 0 { // 偶数
			if col == n-1 {
				// 往下移动一格，准备向下遍历
				row++
			} else if row == 0 {
				// 往右移动一格，准备向下遍历
				col++
			} else {
				// 往上移动
				row--
				col++
			}
		} else { // 奇数
			if row == m-1 {
				// 往右移动一格，准备向上遍历
				col++
			} else if col == 0 {
				// 往下移动一格，准备向上遍历
				row++
			} else {
				// 往下移动
				row++
				col--
			}
		}
	}
	return res
}
