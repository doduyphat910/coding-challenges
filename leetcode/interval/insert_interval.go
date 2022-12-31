package interval

func insert(intervals [][]int, newInterval []int) [][]int {
	var result = make([][]int, 0)
	for i := range intervals {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		}
		if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
			continue
		}
		if intervals[i][0] <= newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] >= newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
	}
	result = append(result, newInterval)

	return result
}
