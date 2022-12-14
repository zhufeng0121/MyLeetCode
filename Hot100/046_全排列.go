package Hot100

// 给定一个不包含重复数字的数组nums，返回所有可能的全排列 (这方法不是很好理解)
func permute(nums []int) [][]int {
	// 处理边界情况
	res := make([][]int, 0)
	backtrack(nums, []int{}, &res)
	return res
}

func backtrack(nums, prev []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, prev)
		return
	}

	for i, v := range nums {
		// 获取除 i 以外的元素
		temp := append(append([]int{}, nums[:i]...), nums[i+1:]...)

		backtrack(temp, append(prev, v), res)
	}
}

func permuteI(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	length := len(nums)
	res := make([][]int, 0)

	// 采用前后元素交换的办法，dfs解题

	exchange(nums, 0, length, &res)
	return res
}

func exchange(nums []int, i, length int, res *[][]int) {

}

// 全排列

func permuteV(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used := make([]bool, len(nums))
	prev := make([]int, 0)
	res := make([][]int, 0)
	generatePermutation(nums, 0, prev, &res, &used)
	return res

}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return

}
