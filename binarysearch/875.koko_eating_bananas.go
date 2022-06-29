package binarysearch

import "math"

// 875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	low, high := 1, 0
	for _, pile := range piles {
		if pile > high {
			high = pile
		}
	}
	k := high
	for low < high {
		speed := low + (high-low)/2
		delta := getTime(piles, speed)
		if delta > h {
			// mid偏小，搜索区间变为 [mid+1, right]
			low = speed + 1
		} else if delta <= h {
			k = speed
			// mid偏大，搜索区间变为 [left, mid]
			high = speed
		}
	}
	return k
}

func getTime(piles []int, speed int) int {
	time := 0
	for _, pile := range piles {
		// 如果一小时能吃掉一堆，则delta + 1
		//if speed > p {
		//	time++
		//	continue
		//}
		//// 如果一小时吃不掉一堆，则计算需要花费几小时
		//time += p / speed
		//if r := p % speed; r != 0 {
		//	time++
		//}
		curTime := math.Ceil(float64(pile / speed))
		time += int(curTime)
	}
	return time
}
