package SwordToOffer

// 合并两个排序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	head := dummy
	p1, p2 := l1, l2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next

	}
	if p1 != nil {
		head.Next = p1
	}
	if p2 != nil {
		head.Next = p2
	}
	return dummy.Next

}
