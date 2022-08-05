package bytedance

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := "0"
	len1, len2 := len(num1), len(num2)

	for i := len1 - 1; i >= 0; i-- {
		curr := ""
		for j := len2 - 2; j >= 0; j-- {
			x, y := 1, 1
			for p := 0; p < len1-i-1; p++ {
				x *= 10
			}
			for q := 0; q < len2-j-1; q++ {
				y *= 10
			}
			tmp := (int(num1[i]-'0') * x) * (int(num2[j]-'0') * y)
			curr = strconv.Itoa(tmp)
		}
		res = addStrings(res, curr)
	}
	return res
}

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
