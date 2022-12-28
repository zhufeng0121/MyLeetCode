package Hot100

//https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/?orderBy=most_relevant

// 岛屿数量  注意参数
func numIslands(grid [][]byte) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				areaI(grid, i, j)
				res++
			}
		}
	}
	return res

}

func areaI(grid [][]byte, r, c int) {
	if !InArea(grid, r, c) {
		return
	}

	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2'

	areaI(grid, r-1, c)
	areaI(grid, r+1, c)
	areaI(grid, r, c-1)
	areaI(grid, r, c+1)
}

func InArea(grid [][]byte, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
