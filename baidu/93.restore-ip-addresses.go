package baidu

import (
	"fmt"
	"strings"
)

var res []string

func restoreIpAddresses(s string) []string {
	length := len(s)
	if length < 4 || length > 12 {
		return nil
	}
	var path []string
	res = nil
	dfs1(s, length, 0, 0, path)
	return res
}

// split 已经分割出多少个ip段
// begin 截取 IP 段的起始位置
// path 记录从根节点到叶子节点的一个路径
// res 记录结果集的变量
func dfs1(s string, length, split, begin int, path []string) {
	if begin == length {
		if split == 4 {
			res = append(res, strings.Join(path, "."))
		}
		return
	}
	// 看到剩下的字符串不够了，就退出, length-begin 表示剩余的还未分割的字符串的位数
	if length-begin < (4-split) || length-begin > 3*(4-split) {
		return
	}
	for i := 0; i < 3; i++ {
		if (begin + i) >= length {
			break
		}
		ipSegment := judgeIfIpSegment(s, begin, begin+i)
		if ipSegment != -1 {
			// 判断是 Ip 段的情况下，才去做截取
			path = append(path, fmt.Sprintf("%d", ipSegment))
			dfs1(s, length, split+1, begin+i+1, path)
			path = path[:len(path)-1]
		}
	}
}

func judgeIfIpSegment(s string, l, r int) int {
	len := r - l + 1
	// 大于 1 位的时候，不能以 0 开头
	if len > 1 && s[l] == '0' {
		return -1
	}
	// 转成 int 类型
	res := 0
	for i := l; i <= r; i++ {
		res = res*10 + int(s[i]-'0')
	}
	if res > 255 {
		return -1
	}
	return res
}
