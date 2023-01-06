package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var result = &TreeNode{Val: root.Val}
	result.Right = invertTree(root.Left)
	result.Left = invertTree(root.Right)
	return result
}
