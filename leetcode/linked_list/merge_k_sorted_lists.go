package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		sortedNode     *ListNode
		headSortedNode *ListNode
		i, minIndex    int
	)

	for len(lists) > 0 {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
			if i > 0 {
				i--
			}
			continue
		}
		if lists[i].Val <= lists[minIndex].Val {
			minIndex = i
		}

		if len(lists)-1 == i {
			var node = &ListNode{
				Val: lists[minIndex].Val,
			}
			if sortedNode == nil {
				sortedNode = node
				headSortedNode = sortedNode
			} else {
				sortedNode.Next = node
				sortedNode = sortedNode.Next
			}
			lists[minIndex] = lists[minIndex].Next

			if lists[minIndex] == nil {
				lists = append(lists[:minIndex], lists[minIndex+1:]...)
			}
			i = 0
			minIndex = 0
			continue
		}
		i++
	}

	return headSortedNode
}
