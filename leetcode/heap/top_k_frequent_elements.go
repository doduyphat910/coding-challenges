package heap

func topKFrequent(nums []int, k int) []int {
	var (
		freqByNumMap = make(map[int]int)
		freqByNumArr = make([][]int, len(nums)+1, len(nums)+1)
		result       = make([]int, 0)
	)

	for i := range nums {
		freq, ok := freqByNumMap[nums[i]]
		if !ok {
			freq = 0
		}
		freq += 1
		freqByNumMap[nums[i]] = freq
	}
	for i := range freqByNumMap {
		if len(freqByNumArr[freqByNumMap[i]]) == 0 {
			freqByNumArr[freqByNumMap[i]] = make([]int, 0)
		}
		freqByNumArr[freqByNumMap[i]] = append(freqByNumArr[freqByNumMap[i]], i)
	}

	for i := len(freqByNumArr) - 1; i >= 0; i-- {
		if len(result) == k {
			break
		}
		if len(freqByNumArr[i]) == 0 {
			continue
		}
		result = append(result, freqByNumArr[i]...)
	}

	return result
}
