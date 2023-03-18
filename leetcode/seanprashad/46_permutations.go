package seanprashad

func permute(nums []int) [][]int {
	var result = make([][]int, 0)

	for i := range nums {
		result = append(result, []int{nums[i]})
	}

	return backtracking(result, nums)
}

func backtracking(result [][]int, nums []int) [][]int {
	if len(result[len(result)-1]) == len(nums) {
		return result
	}

	var temptResult = make([][]int, 0)
	for i := range result {
		for j := range nums {
			if isContains(result[i], nums[j]) {
				continue
			}
			var temp = make([]int, len(result[i]), len(result[i]))
			copy(temp, result[i])
			temp = append(temp, nums[j])
			temptResult = append(temptResult, temp)
		}
	}

	return backtracking(temptResult, nums)
}

func isContains(arr []int, num int) bool {
	for i := range arr {
		if arr[i] == num {
			return true
		}
	}
	return false
}
