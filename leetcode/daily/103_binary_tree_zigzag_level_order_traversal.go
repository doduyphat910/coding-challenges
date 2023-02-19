package daily

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
		if i%2 == 0 {
			list[i] = nodesByLevelMap[i]
			continue
		}
		for j := len(nodesByLevelMap[i]) - 1; j >= 0; j-- {
			list[i][j] = nodesByLevelMap[i][len(nodesByLevelMap[i])-1-j]
		}
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
	nodes = append(nodes, root.Val)
	nodesByLevelMap[level] = nodes
	addNode(root.Right, nodesByLevelMap, level+1)
}
