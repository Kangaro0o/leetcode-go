package didi

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	m, n := len(matrix)-1, len(matrix[0])-1
	i, j := 0, 0
	left, right := matrix[i][j], matrix[m][n]
	for left < right {
		midA, midB := (i+m)/2, (j+n)/2
		if matrix[midA][midB] > target {
			m--
			n--
		} else if matrix[midA][midB] < target {
			i++
			j++
		} else {
			return true
		}
	}
	return false
}
