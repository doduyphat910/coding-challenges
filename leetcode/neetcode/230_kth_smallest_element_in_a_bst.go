package neetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var (
		count  int
		result int
	)
	lnr(root, &count, k, &result)
	return result
}

func lnr(root *TreeNode, count *int, k int, result *int) {
	if root == nil {
		return
	}
	if *result != 0 {
		return
	}
	lnr(root.Left, count, k, result)
	*count += 1
	if *count == k {
		*result = root.Val
	}
	lnr(root.Right, count, k, result)
}
