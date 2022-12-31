package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var list = &ListNode{
		Val: list1.Val,
	}
	if list1.Val > list2.Val {
		list.Val = list2.Val
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}
	first := list

	for list1 != nil || list2 != nil {
		node := &ListNode{}
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}
		list.Next = node
		list = list.Next
	}

	return first
}
