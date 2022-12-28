package main

/**
leetcode 082. Remove Duplicates from Sorted List II

Given the head of a sorted linked list, delete all nodes that have
duplicate numbers, leaving only distinct numbers from the original
list. Return the linked list sorted as well.

*/
//
func deleteDuplicatesI(head *ListNode) *ListNode {
	//先处理特殊情况，即节点数有0或1个的时候
	if head == nil || head.Next == nil {
		return head
	}
	//构造虚拟头节点
	Head := &ListNode{
		Next: head,
	}
	//start 和 end 代表相同节点的区间起始节点和终止节点
	start, end := head, head.Next
	//start节点的前驱节点用于删除
	pre := virtualHead

	//循环退出的条件是 start 和 end 同时到最后一个节点
	for start.Next != nil && end.Next != nil {
		for start.Val != end.Val {
			//如果start和end不相等，双方均向后移动
			start = start.Next
			end = end.Next
			pre = pre.Next

		}
		for start.Val == end.Val {
			//如果相等的话，start不动，end向后移动直到不相等
			end = end.Next
		}
		pre.Next = end
		start = end

	}

	return nil

}
