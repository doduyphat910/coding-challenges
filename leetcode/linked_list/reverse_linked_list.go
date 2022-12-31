package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre = &ListNode{
		Val: head.Val,
	}
	for head.Next != nil {
		var last = &ListNode{
			Val: head.Next.Val,
		}
		last.Next = pre
		pre = last

		head = head.Next
	}

	return pre
}
