package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB

	// curA 和 curB 走 A + B 的长度, 如果有相交的节点，那么一定会在相交的节点相遇，如果没有相交的节点，那么 curA 和 curB 最后
	// 均为 nil, 此时 curA == curB，同样可以退出
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA

}
