package main

type NumMatrix struct {
	sum [][]int
}

// sums[i] 为 matrix[i]的前缀和数组
func Constructor(matrix [][]int) NumMatrix {
	// sum 应该是一个长 len(matrix[0]) + 1,高度为 len(matrix)的二维数组
	sum := make([][]int, len(matrix))

	for i, row := range matrix {
		sum[i] = make([]int, len(row)+1)
		for j, v := range row {
			sum[i][j+1] = sum[i][j] + v
		}
	}

	return NumMatrix{sum: sum}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	sum := 0

	for i := row1; i <= row2; i++ {
		sum += this.sum[i][col2+1] - this.sum[i][col1]
	}

	return sum

}

