package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// func maxDepth(root *TreeNode) int {
// 	var max int
// 	var visitedMap = make(map[*TreeNode]struct{})
// 	visit(root, visitedMap, &max, 0)
// 	return max
// }

// func visit(root *TreeNode, visitedMap map[*TreeNode]struct{}, max *int, curr int) {
// 	if root == nil {
// 		return
// 	}

// 	if _, ok := visitedMap[root]; !ok {
// 		visitedMap[root] = struct{}{}
// 		curr += 1
// 	} else {
// 		curr -= 1
// 	}
// 	if curr > *max {
// 		*max = curr
// 	}
// 	visit(root.Left, visitedMap, max, curr)
// 	visit(root.Right, visitedMap, max, curr)
// }
