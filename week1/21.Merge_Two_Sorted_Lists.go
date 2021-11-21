/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	protect := &ListNode{}
	head := protect
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			fmt.Printf("l1 < l2, l2: %d insert l1: %d\n", l2.Val, l1.Val)
			head.Next = l1
			l1 = l1.Next

		} else if l1.Val >= l2.Val {
			fmt.Printf("l1 >= l2, l1: %d insert l2: %d\n", l1.Val, l2.Val)
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 == nil && l2 != nil {
		head.Next = l2
	} else if l2 == nil && l1 != nil {
		head.Next = l1
	}
	return protect.Next
}

