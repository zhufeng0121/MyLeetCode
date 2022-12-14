package SwordToOffer

// 根据 BST的性质来判断，最后一个元素是根，左边没有比根节点大的，右边没有比根节点小的
func verifyPostorder(postorder []int) bool {
	return helper(postorder, 0, len(postorder)-1)
}

func helper(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	mid := left
	root := postorder[right]
	for postorder[mid] < root {
		mid++
	}
	temp := mid
	for temp < right {
		if postorder[temp] < root {
			return false
		}
		temp++
	}
	return helper(postorder, left, mid-1) && helper(postorder, mid, right-1)

}