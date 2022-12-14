package Hot100

// 采用虚拟头结点来求解
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}

	p1, p2 := list1, list2
	p3 := dummy

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p3.Next = p1
			p1 = p1.Next
		} else {
			p3.Next = p2
			p2 = p2.Next
		}
		p3 = p3.Next
	}
	if p1 != nil {
		p3.Next = p1
	}
	if p2 != nil {
		p3.Next = p2
	}

	return dummy.Next

}

// 也可以用递归来做，感觉递归更加优雅
func mergeTwoListsI(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsI(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsI(list1, list2.Next)
		return list2
	}
}
