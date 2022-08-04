package bytedance

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	row, col := 0, 0
	p, q := 1, 0
	for i := 0; i < len(s); i++ {
		res[row] = append(res[row], s[i])
		if row == 0 {
			p = 1
			q = 0
		} else if row == numRows-1 {
			p = -1
			q = 1
		}
		row = row + p
		col = col + q
	}
	var resStr string
	for _, resByte := range res {
		for _, r := range resByte {
			resStr += string(r)
		}
	}
	return resStr
}
