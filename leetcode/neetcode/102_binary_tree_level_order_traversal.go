/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		q       simpleQ
		results = make([][]int, 0)
		temp    = make([]*TreeNode, 0)
		result  = make([]int, 0)
	)
	q.push(root)
	for !q.isEmpty() {
		node := q.pop()
		result = append(result, node.Val)
		if node.Left != nil {
			temp = append(temp, node.Left)
		}
		if node.Right != nil {
			temp = append(temp, node.Right)
		}

		if q.isEmpty() {
			results = append(results, result)
			result = make([]int, 0)
			q.setQ(temp)
			temp = make([]*TreeNode, 0)
		}
	}

	return results
}

type simpleQ struct {
	arr []*TreeNode
}

func (q *simpleQ) setQ(arr []*TreeNode) {
	q.arr = arr
}

func (q *simpleQ) push(val *TreeNode) {
	q.arr = append(q.arr, val)
}

func (q *simpleQ) pop() *TreeNode {
	val := q.arr[0]
	if len(q.arr) > 0 {
		q.arr = q.arr[1:]
	}
	return val
}

func (q *simpleQ) isEmpty() bool { return len(q.arr) == 0 }