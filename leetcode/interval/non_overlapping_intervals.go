package interval

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var (
		result int
		item   = intervals[0]
	)
	for i := 1; i < len(intervals); i++ {
		if item[1] > intervals[i][0] && item[1] <= intervals[i][1] {
			result += 1
			continue
		} else if intervals[i][1] < item[1] {
			result += 1
			item = intervals[i]
			continue
		}
		item = intervals[i]
	}
	return result
}
