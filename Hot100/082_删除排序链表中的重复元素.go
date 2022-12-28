package Hot100

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 创建一个虚拟头结点
	dummy := &ListNode{
		Next: head,
		Val:  -1010,
	}

	// 删除节点需要该节点的前驱节点

	pre := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			val := cur.Val
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}
