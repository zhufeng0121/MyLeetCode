package Hot100

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完全二叉树 可能是满二叉树，但 满二叉树 不一定是完全二叉树, 根据左右子树 判断是否是满二叉树
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	hl, hr := 0, 0

	// 求左子树的高度
	for l != nil {
		hl++
		l = l.Left

	}

	for r != nil {
		hr++
		r = r.Right

	}

	if hl == hr {
		return int(math.Pow(2, float64(hl))) - 1
	}

	// 当做满二叉树来计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}
