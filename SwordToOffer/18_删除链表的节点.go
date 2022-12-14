package SwordToOffer

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	dummy := &ListNode{Next: head}
	cur := head
	pre := dummy

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		} else {
			pre = pre.Next
			cur = cur.Next

		}
	}
	return dummy.Next

}
