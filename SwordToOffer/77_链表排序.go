package SwordToOffer

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 利用快慢指针找出中点，要找出中点的前一个节点
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 对右半部分进行归并排序
	right := sortList(slow.Next)
	// 断开
	slow.Next = nil
	left := sortList(head)

	return mergeList(left, right)

}

// 先实现一个合并两个排序链表
func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	// 虚拟节点
	dummy := &ListNode{Val: -1}
	cur := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return dummy.Next
}
