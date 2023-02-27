package seanprashad

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var nodesByLevelMap = make(map[int][]int, 0)
	addNode(root, nodesByLevelMap, 0)

	var list = make([][]int, len(nodesByLevelMap), len(nodesByLevelMap))
	for i := 0; i < len(nodesByLevelMap); i++ {
		list[i] = make([]int, len(nodesByLevelMap[i]), len(nodesByLevelMap[i]))
		list[i] = nodesByLevelMap[i]
	}
	return list
}

func addNode(root *TreeNode, nodesByLevelMap map[int][]int, level int) {
	if root == nil {
		return
	}
	addNode(root.Left, nodesByLevelMap, level+1)
	nodes, ok := nodesByLevelMap[level]
	if !ok {
		nodes = make([]int, 0)
	}
	if level%2 == 0 {
		nodes = append(nodes, root.Val)
	} else {
		nodes = append([]int{root.Val}, nodes...)
	}
	nodesByLevelMap[level] = nodes
	addNode(root.Right, nodesByLevelMap, level+1)
}
