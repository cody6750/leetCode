package questions

type ListNode struct {
	Val  int
	Next *ListNode
}

// [ 1 2 3] [ 3 2 1]
// [1 1 2 2 3 3]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var finalListNode ListNode
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l2
	}

	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			finalListNode.Next = l2
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			finalListNode.Next = l1
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			finalListNode.Next = l2
			l2 = l2.Next
		} else if l1.Val > l2.Val {
			finalListNode.Next = l1
			l1 = l2.Next
		}
	}

	return &finalListNode

}
