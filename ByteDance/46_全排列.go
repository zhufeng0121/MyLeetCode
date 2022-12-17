package ByteDance

// 计算全排列的数量
func permute(nums []int) [][]int {
	res := [][]int{}

	used := make([]bool, len(nums))

	var dfs func(tmp []int, used []bool)

	dfs = func(tmp []int, used []bool) {
		if len(tmp) == len(nums) {
			c := make([]int, len(nums))
			copy(c, tmp)
			res = append(res, c)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == false {
				tmp = append(tmp, nums[i])
				used[i] = true
				dfs(tmp, used)
				tmp = tmp[:len(tmp)-1]
				used[i] = false
			}
		}

	}

	dfs([]int{}, used)

	return res

}
