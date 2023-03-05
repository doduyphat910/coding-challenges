package seanprashad

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var max = math.MinInt
	lnr(root, &max)
	return max
}

func lnr(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := lnr(root.Left, max)
	right := lnr(root.Right, max)

	if *max < left+right {
		*max = left + right
	}

	if left > right {
		return left + 1
	}
	return right + 1
}
