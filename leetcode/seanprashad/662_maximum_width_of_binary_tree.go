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
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q Queue
	var max = math.MinInt
	rootNode := NodeMap{
		Node:  root,
		Index: 1,
	}
	q.Push(rootNode)
	for !q.IsEmpty() {
		qLength := q.Length()
		left := q.arr[0]
		right := q.arr[qLength-1]
		if right.Index-left.Index+1 > max {
			max = right.Index - left.Index + 1
		}
		for i := 0; i < qLength; i++ {
			node := q.Pop()
			if node.Node.Left != nil {
				newNode := NodeMap{
					Node:  node.Node.Left,
					Index: node.Index * 2,
				}
				q.Push(newNode)
			}
			if node.Node.Right != nil {
				newNode := NodeMap{
					Node:  node.Node.Right,
					Index: node.Index*2 + 1,
				}
				q.Push(newNode)
			}
		}
	}

	return max
}

type NodeMap struct {
	Node  *TreeNode
	Index int
}

// Queue
type Queue struct {
	arr []NodeMap
}

func (q *Queue) Length() int {
	return len(q.arr)
}

func (q *Queue) Pop() NodeMap {
	first := q.arr[0]
	q.arr = q.arr[1:]
	return first
}

func (q *Queue) Push(item NodeMap) []NodeMap {
	q.arr = append(q.arr, item)
	return q.arr
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}
