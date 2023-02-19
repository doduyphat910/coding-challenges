package daily

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	var (
		arrNode = make([]int, 0)
		minDiff = math.MaxInt
	)
	lnr(root, &arrNode)
	for i := 1; i < len(arrNode); i++ {
		diff := arrNode[i] - arrNode[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func lnr(root *TreeNode, arrNode *[]int) []int {
	if root == nil {
		return *arrNode
	}

	lnr(root.Left, arrNode)
	*arrNode = append(*arrNode, root.Val)
	lnr(root.Right, arrNode)

	return *arrNode
}

//Option 2:
// func minDiffInBST(root *TreeNode) int {
//     var min = math.MaxInt
//     var prevVal = -1
//     lnr(root, &prevVal, &min)
//     return min
// }

// func lnr(root *TreeNode, prevVal *int, min *int) {
//     if root == nil {
//         return
//     }

//     lnr(root.Left, prevVal, min)
//     if *prevVal != -1 && root.Val - *prevVal < *min {
//         *min = root.Val - *prevVal
//     }
//     *prevVal = root.Val
//     lnr(root.Right, prevVal, min)
// }
