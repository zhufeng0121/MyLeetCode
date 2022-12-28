package Hot100

// https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/?orderBy=most_relevant
func maxAreaOfIsland(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 对所有为1的进行dfs
			if grid[i][j] == 1 {
				res = maxInt(res, area(grid, i, j))
			}
		}
	}
	return res

}

func area(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 0
	}

	if grid[r][c] != 1 {
		return 0
	}
	// 进行标记，防止回头
	grid[r][c] = 2

	return 1 + area(grid, r-1, c) + area(grid, r+1, c) + area(grid, r, c-1) + area(grid, r, c+1)
}

// 判断 r,c 是否在 grid格子之中
func inArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
