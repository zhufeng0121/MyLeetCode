package main

func checkXMatrix(grid [][]int) bool {
	m := len(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if (i == j || i+j == m-1) && grid[i][j] == 0 {
				return false
			} else if (i != j && i+j != m-1) && grid[i][j] != 0 {
				return false

			}

		}
	}
	return true

}