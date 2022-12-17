package Hot100

func combine(n int, k int) [][]int {
	res := [][]int{}

	track := make([]int, 0)

	var dfs func(n int, k int, start int, track []int)

	dfs = func(n int, k int, start int, track []int) {

		if k == len(track) {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := start; i <= n; i++ {
			track = append(track, i)
			dfs(n, k, i+1, track)
			track = track[:len(track)-1]
		}

	}

	dfs(n, k, 1, track)
	return res

}
