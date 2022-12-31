package interval

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var results = [][]int{intervals[0]}
	for i := range intervals {
		if intervals[i][0] > results[len(results)-1][1] {
			results = append(results, intervals[i])
			continue
		}
		if intervals[i][0] <= results[len(results)-1][0] {
			results[len(results)-1][0] = intervals[i][0]
		}
		if intervals[i][1] >= results[len(results)-1][1] {
			results[len(results)-1][1] = intervals[i][1]
		}
	}

	return results
}
