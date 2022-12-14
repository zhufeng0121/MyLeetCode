package SwordToOffer

// fast 先走K步，然后 slow 和 fast 一起走
func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	for k != 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow

}
