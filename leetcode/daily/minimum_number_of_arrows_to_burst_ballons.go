package daily

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	var end = points[0][1]
	var arrow = 1
	for i := 1; i < len(points); i++ {
		if end < points[i][0] {
			end = points[i][1]
			arrow += 1
		}
	}

	return arrow
}
