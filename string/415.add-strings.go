package string

import "strconv"

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0 // 进位
	ans := ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}
