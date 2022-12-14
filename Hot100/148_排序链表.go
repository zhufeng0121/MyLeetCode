package Hot100

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]int, 0)

	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	sort.Ints(arr)

	newHead := &ListNode{
		Next: nil,
		Val:  arr[0],
	}
	cur = newHead
	for i := 1; i < len(arr); i++ {
		newHead.Next = &ListNode{
			Next: nil,
			Val:  arr[i],
		}
		newHead = newHead.Next
	}
	return cur

}

// 归并排序
func sortListI(head *ListNode) *ListNode {
	return mergeSort(head)

}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找出中位点
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//对右半部分进行归并排序
	right := mergeSort(slow.Next)
	slow.Next = nil
	//对左半部分进行归并排序
	left := mergeSort(head)

	return mergeList(left, right)
}

func mergeList(left, right *ListNode) *ListNode {
	//临时头结点
	tmpHead := &ListNode{}

	p := tmpHead

	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}

	if left == nil {
		p.Next = right
	} else {
		p.Next = left
	}
	return tmpHead.Next

}

// 链表快排
