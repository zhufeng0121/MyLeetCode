package SwordToOffer

func subsets(nums []int) [][]int {
	res := [][]int{}

	track := make([]int, 0)

	var dfs func(start int, track []int)

	dfs = func(start int, track []int) {

		c := make([]int, len(track))
		copy(c, track)

		res = append(res, c)

		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			dfs(i+1, track)
			track = track[:len(track)-1]
		}

	}

	dfs(0, track)
	return res
}
