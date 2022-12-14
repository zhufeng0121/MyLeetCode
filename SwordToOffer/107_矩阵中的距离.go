package SwordToOffer

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}


	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = mat[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] != 0 {
				// 寻找四周离他最近的1
				if i == 0 && j== 0{
					res[i][j] = 
				}
				res[i][j] = maxUtil(res[])
			}
		}
	}


}
