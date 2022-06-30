package binarysearch

// 1011. 在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	// 计算船的最大运载重量
	low, high := 0, 0
	for _, weight := range weights {
		if weight > low {
			low = weight
		}
		high += weight
	}
	k := high
	for low < high {
		load := low + (high-low)/2
		needDays := getDays(weights, load)
		if needDays > days {
			// 超过了目标天数,则说明运载能力低，运载能力区间变为 [load+1, high]
			low = load + 1
		} else if needDays <= days {
			// 低于目标天数，说明船只运载能力偏高，降低运载能力区间 [low, load]
			k = load
			high = load
		}
	}
	return k
}

func getDays(weights []int, load int) int {
	days := 0
	w := 0 // 连续装载包裹的和
	for i := 0; i < len(weights); i++ {
		w += weights[i]
		if w > load {
			days++
			w = weights[i]
		}
	}
	if w <= load {
		days++
	}
	return days
}
