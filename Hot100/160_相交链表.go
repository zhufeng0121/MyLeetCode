package Hot100

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB

	// 1.程序一定会退出，无论A 和 B 是否相交
	for curA != curB {

		if curA == nil {
			// 应该赋值的是headB, 而不是curB 此处一定要小心
			curA = headB
		} else {
			// 这么写的原因是通过对curA 和 curB 重新赋值，来替代前进一步
			curA = curA.Next
		}

		if curB == nil {
			// 同上
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}

