package Hot100

// 采用 合并两个升序链表的方法试一试
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	newLists := make([]*ListNode, 0)
	if len(lists)%2 == 0 {
		for i := 0; i < len(lists); i += 2 {
			newHead := mergeTwo(lists[i], lists[i+1])
			newLists = append(newLists, newHead)
		}

	} else {
		for i := 0; i < len(lists)-1; i += 2 {
			newHead := mergeTwo(lists[i], lists[i+1])
			newLists = append(newLists, newHead)

		}
		newLists = append(newLists, mergeTwo(lists[len(lists)-1], nil))
	}

	head := mergeKLists(newLists)

	return head

}

func mergeTwo(list1 *ListNode, list2 *ListNode) *ListNode {
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