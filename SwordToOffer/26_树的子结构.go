package SwordToOffer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	} else {
		return helper(A, B) || isSubStructure(A.Right, B) || isSubStructure(A.Left, B)
	}

}

func helper(A *TreeNode, B *TreeNode) bool {
	//
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return helper(A.Left, B.Left) && helper(A.Right, B.Right)

}
