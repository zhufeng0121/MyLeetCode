package Hot100

// dan
type ListNode struct {
	Val  int
	Next *ListNode
}

// 两个链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0          // 进位
	head := &ListNode{} //虚拟头结点
	current := head

	for l1 != nil && l2 != nil {
		value := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		current.Next = node
		current = current.Next

		//l1 和 l2 向下走
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		value := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10

		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		current.Next = node
		current = current.Next

		// l1 向下走
		l1 = l1.Next

	}

	for l2 != nil {
		value := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10

		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		current.Next = node
		current = current.Next

		l2 = l2.Next

	}
	if carry == 1 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}
		current.Next = node
		current = current.Next

	}

	return head.Next

}
