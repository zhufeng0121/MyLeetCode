package SwordToOffer

// 有一种比较比较容易得解决办法是，获取BST的中序遍历数组，然后利用双指针法

// 题解： https://leetcode.cn/problems/opLdQZ/solutions/1350500/er-cha-sou-suo-shu-zhong-liang-ge-jie-di-bqci/
func findTarget(root *TreeNode, k int) bool {
	res := []int{}

	inorder(root, &res)

	start, end := 0, len(res)-1

	for start < end {
		if res[start]+res[end] == k {
			return true
		} else if res[start]+res[end] > k {
			end--
		} else if res[start]+res[end] < k {
			start++
		}
	}
	return false

}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root, res)
	*res = append(*res, root.Val)
	inorder(root, res)

}

func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]struct{})
	var dfs func(*TreeNode) bool

	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}

		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)

	}
	return dfs(root)

}
