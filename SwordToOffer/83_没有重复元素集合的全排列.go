package SwordToOffer

func permute(nums []int) [][]int {
	res := [][]int{}

	visited := map[int]bool{}

	var dfs func(path []int)

	dfs = func(path []int) {
		// 退出条件
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return

		}
		// 每一次遍历选择一个数据
		for i := 0; i < len(nums); i++ {
			n := nums[i]
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1] // 撤回选择
			visited[n] = false

		}
	}
	dfs([]int{})
	return res

}
