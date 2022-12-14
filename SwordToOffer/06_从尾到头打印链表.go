package SwordToOffer

// 这个问题的解法有很多，怎么解都行
func reversePrint(head *ListNode) []int {
	newHead := reverse(head)
	res := make([]int, 0)

	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res

}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
