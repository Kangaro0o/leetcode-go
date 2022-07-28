package pdd

// numIslands 200. 岛屿数量
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	nums := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' { // 发现了岛屿
				nums++
				// 淹没此岛屿，目的是可以避免记录 visited 数组
				dfs(grid, i, j)
			}
		}
	}
	return nums
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	// 淹没此岛屿
	grid[i][j] = '0'
	dfs(grid, i-1, j) // 上
	dfs(grid, i+1, j) // 下
	dfs(grid, i, j-1) // 左
	dfs(grid, i, j+1) // 右
}
