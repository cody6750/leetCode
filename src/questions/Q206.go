package questions

type LinkedList struct {
	Val  int
	Next *LinkedList
}

var finalNode *LinkedList

func reverseList(head *LinkedList) *LinkedList {
	if head == nil {
		return head
	}
	if head.Next != nil {
		//reverse
		finalNode = reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
	} else {
		finalNode = head
	}
	return finalNode
}
