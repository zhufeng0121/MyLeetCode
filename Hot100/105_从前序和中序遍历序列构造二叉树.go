package Hot100

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	root := &TreeNode{
		Val: rootVal,
	}
	index := findElem(inorder, rootVal)

	leftArr := inorder[:index]
	rightArr := inorder[index+1:]

	root.Left = buildTree(preorder[1:len(leftArr)+1], leftArr)
	root.Right = buildTree(preorder[len(leftArr)+1:], rightArr)

	return root

}

func findElem(arr []int, target int) (index int) {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return

}
