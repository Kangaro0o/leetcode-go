package ali

// twoSum 1. 两数之和
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int, len(nums))
	for idx, num := range nums {
		if val, ok := memo[num]; ok {
			return []int{val, idx}
		} else {
			memo[target-num] = idx
		}
	}
	return nil
}
