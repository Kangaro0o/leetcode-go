package baidu

func permute(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	var res [][]int
	visited := make([]bool, length)
	var path []int
	dfs(nums, length, 0, path, visited, res)
	return res
}

func dfs(nums []int, length, depth int, path []int, visited []bool, res [][]int) {
	if depth == length {
		tmp := make([]int, len(path))
		for i := 0; i < len(path); i++ {
			tmp[i] = path[i]
		}
		res = append(res, tmp)
		return
	}

	for i := 0; i < length; i++ {
		if !visited[i] {
			path = append(path, nums[i])
			visited[i] = true

			dfs(nums, length, depth+1, path, visited, res)

			// 回溯
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
}
