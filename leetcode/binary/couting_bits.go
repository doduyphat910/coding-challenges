package binary

// brute-force
func countBits(n int) []int {
	var result = make([]int, 0)
	var i int

	for i <= n {
		var total1 int
		a := i
		for a != 0 {
			if (a & 1) != 0 {
				total1 += 1
			}
			a = a >> 1
		}
		result = append(result, total1)
		i++
	}

	return result
}

// optimize O(N)
func countBits(n int) []int {
	ret := []int{0}
	for i := 1; i <= n; i++ {
		info := 0
		if i&1 == 1 {
			info = ret[len(ret)-1] + 1
		} else {
			info = ret[len(ret)/2]
		}
		ret = append(ret, info)
	}
	return ret
}
