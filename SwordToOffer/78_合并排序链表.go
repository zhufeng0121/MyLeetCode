package SwordToOffer

// 合并K个排序链表
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length >> 1
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return merge2Lists(left, right)

}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = merge2Lists(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge2Lists(list1, list2.Next)
		return list2
	}

}
