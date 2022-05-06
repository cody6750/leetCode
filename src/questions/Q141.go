/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	ptr := head.Next
	for head != nil && ptr != nil {
		if ptr == head {
			return true
		}
		if ptr.Next == nil {
			return false
		}
		ptr = ptr.Next.Next
		head = head.Next
	}
	return false
}