/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root.Val == subRoot.Val {
		if isSameTree(root, subRoot) {
			return true
		}
	}
	if root.Left != nil {
		if isSubtree(root.Left, subRoot) {
			return true
		}
	}
	if root.Right != nil {
		if isSubtree(root.Right, subRoot) {
			return true
		}
	}
	return false
}

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil ||
		root != nil && subRoot == nil ||
		root != nil && subRoot != nil && root.Val != subRoot.Val {
		return false
	}
	if (root.Left != nil || subRoot.Left != nil) && !isSameTree(root.Left, subRoot.Left) {
		return false
	}
	if (root.Right != nil || subRoot.Right != nil) && !isSameTree(root.Right, subRoot.Right) {
		return false
	}
	return true
}