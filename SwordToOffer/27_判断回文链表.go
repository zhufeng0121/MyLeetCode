package SwordToOffer

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	//快慢指针寻找终点
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 如果 fast ！= nil,链表的长度为奇数，将slow 向后挪一位
	if fast != nil {
		slow = slow.Next
	}

	fast = head
	slow = reverseLListII(slow)

	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true

}

func reverseLListII(head *ListNode) *ListNode {
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
