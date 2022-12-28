package Hot100

func islandPerimeter(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res = areaII(grid, i, j)
			}

		}
	}
	return res

}

func areaII(grid [][]int, r, c int) int {

	// 到了边上 +1
	if !InArea(grid, r, c) {
		return 1
	}
	// 到了海上 +1
	if grid[r][c] == 0 {
		return 1
	}

	// 遍历已经遍历过的
	if grid[r][c] != 1 {
		return 0
	}

	// 标记已经遍历过
	grid[r][c] = 2

	return areaII(grid, r+1, c) + areaII(grid, r-1, c) + areaII(grid, r, c-1) + areaII(grid, r, c+1)

}

func InArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}