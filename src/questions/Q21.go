/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	"fmt"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	final := &ListNode{}
	head := final
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				final.Val = list1.Val
				list1 = list1.Next
			} else {
				final.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 != nil {
			final.Val = list1.Val
			list1 = list1.Next
		} else if list2 != nil {
			final.Val = list2.Val
			list2 = list2.Next
		}
		fmt.Print(final)
		if list1 == nil && list2 == nil {
			break
		}
		final.Next = &ListNode{}
		final = final.Next
	}

	return head
}

// [] []
// [1] []
// [] [2]
// [1 1 1 ] [2 2 3 ]
// [-1]  [5]
// [5] [-1]