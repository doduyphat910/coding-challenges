package array

// Brute Force
func maxArea(height []int) int {
	var max int
	for i := range height {
		for j := len(height) - 1; j >= 0; j-- {
			if j <= i {
				break
			}
			var result int
			if height[i] < height[j] {
				result = height[i] * (j - i)
			} else {
				result = height[j] * (j - i)
			}
			if result > max {
				max = result
			}
		}
	}

	return max
}

// another approach
func maxArea(height []int) int {
	var (
		i   = 0
		j   = len(height) - 1
		max int
	)

	for i < j {
		var result int
		if height[i] < height[j] {
			result = height[i] * (j - i)
			i++
		} else {
			result = height[j] * (j - i)
			j--
		}
		if result > max {
			max = result
		}
	}

	return max
}
