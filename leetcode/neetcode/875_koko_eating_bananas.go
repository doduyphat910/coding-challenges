package neetcode

import "math"

func minEatingSpeed(piles []int, h int) int {
	var max int
	for i := range piles {
		if piles[i] > max {
			max = piles[i]
		}
	}
	var (
		l   = 1
		r   = max
		min = math.MaxInt
	)
	for l <= r {
		var (
			count int
			mid   = (l + r) / 2
		)
		for i := range piles {
			count += int(math.Ceil(float64(piles[i]) / float64(mid)))
		}
		if count > h {
			l = mid + 1
		}
		if count <= h {
			r = mid - 1
			if mid < min {
				min = mid
			}
		}
	}

	return min
}
