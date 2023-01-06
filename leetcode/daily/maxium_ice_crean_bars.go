package daily

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})
	var max int
	for i := range costs {
		if coins <= 0 {
			break
		}
		if coins-costs[i] >= 0 {
			coins -= costs[i]
			max += 1
			continue
		}
		i = 0
	}

	return max
}
