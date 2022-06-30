package slidingwindow

// 239. 滑动窗口最大值
func maxSlidingWindow1(nums []int, k int) []int {
	max := nums[0]
	maxIndex := 0
	left, right := 0, 0
	var res []int // 记录结果
	for right < len(nums) {
		c := nums[right]
		if maxIndex < left { // 最大值不在窗口范围内
			// 重新寻找窗口内最大值
			max, maxIndex = maxInt(nums[left:right+1], left)
		}
		if c > max {
			max = c
			maxIndex = right
		}
		right++
		if (right - left) == k {
			res = append(res, max)
			left++
		}
	}
	return res
}

func maxInt(nums []int, offset int) (int, int) {
	max := nums[0]
	maxIndex := 0
	for i, n := range nums {
		if n > max {
			max = n
			maxIndex = i
		}
	}
	return max, maxIndex + offset
}

// 1. 维护一个单调队列
// 2。 使用单调队列的特性（队首是最大的元素）来方便取出窗口范围内的最大值
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMonotonicQueue()
	var res []int
	for i := 0; i < len(nums); i++ {
		if i+1 < k {
			queue.push(nums[i])
		} else {
			queue.push(nums[i])
			max := queue.max()
			res = append(res, max)
			// 益处滑动窗口最左侧数字
			queue.pop(nums[i+1-k])
		}
	}
	return res
}

// MonotonicQueue 单调队列
type MonotonicQueue struct {
	data []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		data: make([]int, 0),
	}
}

func (m *MonotonicQueue) push(n int) {
	if len(m.data) > 0 && m.data[0] < n { // 队列中所有元素都比 n 小
		m.data = []int{}
	}
	for len(m.data) > 0 && m.data[len(m.data)-1] < n { // 队列中从队尾开始存在元素比 n 小
		m.data = m.data[:len(m.data)-1] // 删除最后一个元素
	}
	m.data = append(m.data, n)
}

func (m *MonotonicQueue) pop(n int) {
	if m.data[0] == n { // 判断滑动窗口最左边元素t是不是单调队列最大值，若是弹出，否则什么都不用做
		m.data = m.data[1:]
	}
}

func (m *MonotonicQueue) max() int {
	return m.data[0]
}
