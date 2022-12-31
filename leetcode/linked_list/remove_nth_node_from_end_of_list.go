package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		totalNode  int
		headTotal  = head
		headRemove = head
		current    int
	)
	for headTotal != nil {
		totalNode += 1
		headTotal = headTotal.Next
	}

	var first = headRemove
	for headRemove != nil {
		current += 1
		preRemoveIndex := totalNode - n
		if preRemoveIndex <= 0 {
			next := headRemove.Next
			return next
		}
		if current == preRemoveIndex {
			headRemove.Next = headRemove.Next.Next
		}
		headRemove = headRemove.Next
	}

	return first
}
