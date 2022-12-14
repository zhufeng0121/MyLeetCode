package SwordToOffer

// 第一种简单做法，将两个链表都反转，进行调用相加函数，再反转回来
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head1, head2 := reverseListForAdd(l1), reverseListForAdd(l2)
	head3 := AddTwoListReverse(head1, head2)
	return reverseListForAdd(head3)

}

func reverseListForAdd(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev

}

func AddTwoListReverse(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	// 创建一个虚拟头结点
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		value := (carry + l1.Val + l2.Val) % 10
		carry = (carry + l1.Val + l2.Val) / 10
		next := &ListNode{Val: value}
		cur.Next = next
		cur = cur.Next
		// l1 和 l2 向下走
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		value := (carry + l1.Val) % 10
		carry = (carry + l1.Val) / 10
		next := &ListNode{Val: value}
		cur.Next = next
		cur = cur.Next
		// l1 向下走
		l1 = l1.Next

	}
	for l2 != nil {
		value := (carry + l2.Val) % 10
		carry = (carry + l2.Val) / 10
		next := &ListNode{Val: value}
		cur.Next = next
		cur = cur.Next

		// l2 向下走
		l2 = l2.Next
	}

	// 还需要判断一下进位

	if carry == 1 {
		next := &ListNode{Val: carry}
		cur.Next = next
		cur = cur.Next

	}

	return dummy.Next

}
