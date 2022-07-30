package pdd

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Sort(Interval(intervals))
	var res [][]int
	x, y := intervals[0][0], intervals[0][1]
	res = append(res, []int{x, y})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= y { // 区间有重叠
			x = min(intervals[i][0], x)
			y = max(intervals[i][1], y)
			res = res[:len(res)-1]
			res = append(res, []int{x, y})
		} else {
			// 不重叠
			x, y = intervals[i][0], intervals[i][1]
			res = append(res, []int{x, y})
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Interval [][]int

func (i Interval) Len() int {
	return len(i)
}

func (i Interval) Less(p, q int) bool {
	return i[p][0] < i[q][0]
}

func (i Interval) Swap(p, q int) {
	i[p], i[q] = i[q], i[p]
}
