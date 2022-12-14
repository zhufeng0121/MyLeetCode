package main

func diagonalSum(mat [][]int) int {
	m := len(mat)
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i == j || i+j == m-1 {
				sum += mat[i][j]
			}
		}
	}
	return sum

}
