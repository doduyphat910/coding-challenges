package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func hasCycle(head *ListNode) bool {
//     var nodeMap = make(map[*ListNode]struct{})

//     for head != nil {
//         if _, ok := nodeMap[head]; !ok {
//             nodeMap[head] = struct{}{}
//             head = head.Next
//             continue
//         }
//         return true
//     }
//     return false
// }

func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
