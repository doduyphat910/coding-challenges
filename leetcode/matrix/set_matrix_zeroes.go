package matrix

func setZeroes(matrix [][]int) {
	var (
		rowZero    = false
		columnZero = false
	)
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				if i == 0 && !rowZero {
					rowZero = true
				}
				if j == 0 && !columnZero {
					columnZero = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if columnZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
