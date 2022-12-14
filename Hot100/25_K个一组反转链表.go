package Hot100

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 反转前 k 个元素
	newHead := reverseRange(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead

}

func reverseRange(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a

	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre

}