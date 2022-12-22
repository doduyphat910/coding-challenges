package graph

func pacificAtlantic(heights [][]int) [][]int {
	pacMatrix := createMatrix(heights)
	oceMatrix := createMatrix(heights)

	for i := range heights {
		dfs(heights, i, 0, heights[i][0], pacMatrix)
		dfs(heights, i, len(heights[0])-1, heights[i][len(heights[0])-1], oceMatrix)
	}
	for j := range heights[0] {
		dfs(heights, 0, j, heights[0][j], pacMatrix)
		dfs(heights, len(heights)-1, j, heights[len(heights)-1][j], oceMatrix)
	}

	var result = make([][]int, 0)
	for i := range heights {
		for j := range heights[i] {
			if pacMatrix[i][j] && oceMatrix[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func createMatrix(heights [][]int) [][]bool {
	var result = make([][]bool, len(heights))
	for i := range heights {
		result[i] = make([]bool, len(heights[i]))
	}

	return result
}

func dfs(heights [][]int, i, j int, prev int, visitted [][]bool) {
	if i < 0 || i > len(heights)-1 || j < 0 || j > len(heights[0])-1 {
		return
	}
	if visitted[i][j] || heights[i][j] < prev {
		return
	}

	visitted[i][j] = true
	dfs(heights, i, j+1, heights[i][j], visitted)
	dfs(heights, i, j-1, heights[i][j], visitted)
	dfs(heights, i+1, j, heights[i][j], visitted)
	dfs(heights, i-1, j, heights[i][j], visitted)
}
