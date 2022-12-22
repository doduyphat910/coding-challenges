package matrix

func spiralOrder(matrix [][]int) []int {
	var (
		top, right = 0, len(matrix[0]) - 1
		left, bot  = 0, len(matrix) - 1
		output     = make([]int, 0)
	)

	for left <= right && top <= bot {
		for i := top; i <= right; i++ {
			output = append(output, matrix[top][i])
		}
		top++
		for i := top; i <= bot; i++ {
			output = append(output, matrix[i][right])
		}
		right--
		if left > right || top > bot {
			break
		}
		for i := right; i >= left; i-- {
			output = append(output, matrix[bot][i])
		}
		bot--
		for i := bot; i >= top; i-- {
			output = append(output, matrix[i][left])
		}
		left++
	}

	return output
}
