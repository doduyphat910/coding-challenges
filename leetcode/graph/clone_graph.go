package graph

func (node *Node) setValue(val int) {
	node.Val = val
}

func cloneGraph(node *Node) *Node {
	var (
		traveledSet = make(map[int]*Node)
	)

	if node == nil {
		return node
	}
	return copyGraph(node, traveledSet)
}

func copyGraph(srcNode *Node, traveledSet map[int]*Node) (dstNode *Node) {
	dstNode, ok := traveledSet[srcNode.Val]
	if ok {
		return dstNode
	}

	dstNode = &Node{
		Val:       srcNode.Val,
		Neighbors: make([]*Node, len(srcNode.Neighbors)),
	}
	traveledSet[srcNode.Val] = dstNode

	for i := range srcNode.Neighbors {
		dstNode.Neighbors[i] = copyGraph(srcNode.Neighbors[i], traveledSet)
	}

	return dstNode
}
