package other

// intervalIntersection 986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	p1, p2 := 0, 0
	var res [][]int
	for p1 < len(firstList) && p2 < len(secondList) {
		a := firstList[p1]
		b := secondList[p2]
		if a[1] < b[1] {
			if intersection := calIntersection(a, b); intersection != nil {
				res = append(res, intersection)
			}
			p1++
		} else {
			if intersection := calIntersection(b, a); intersection != nil {
				res = append(res, intersection)
			}
			p2++
		}
	}
	return res
}

func calIntersection(a, b []int) []int {
	x1, y1 := a[0], a[1]
	x2, _ := b[0], b[1]
	if x2 > y1 {
		return nil
	}
	if x2 >= x1 {
		return []int{x2, y1}
	}
	// x2 < x1
	return []int{x1, y1}
}
