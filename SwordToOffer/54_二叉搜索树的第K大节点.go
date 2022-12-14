package SwordToOffer

// BST 中序遍历是一个有序数组
func kthLargest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	helperV(root, &arr)
	return arr[len(arr)-k]

}

func helperV(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		helperV(root.Left, arr)
	}
	*arr = append(*arr, root.Val)
	if root.Right != nil {
		helperV(root.Right, arr)
	}

}

// 当遍历到 第K大的数时，就可以停止遍历了
func kthLargestI(root *TreeNode, k int) int {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)

	}

	dfs(root)

	return nums[len(nums)-k]
}

func kthLargestII(root *TreeNode, k int) int {
	ans := 0
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			ans = root.Val
		}
		dfs(root.Left)
	}

	dfs(root)
	return ans
}