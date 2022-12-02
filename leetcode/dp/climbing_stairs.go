package dp

func climbStairs(n int) int {
	// if n == 1 || n == 2{
	//     return n
	// }

	//     var n1 = 1
	//     var n2 = 2
	//     var total int
	//     for i := 3; i <=n; i++ {
	//         total = (n1 + n2)

	//         n1 = n2
	//         n2 = total
	//     }

	//     return total

	cache := make([]int, 46)
	cache[1] = 1
	cache[2] = 2
	return climbStairsCache(n, cache)
}

func climbStairsCache(n int, cache []int) int {
	if cache[n] != 0 {
		return cache[n]
	}

	cache[n] = climbStairsCache(n-1, cache) + climbStairsCache(n-2, cache)
	return cache[n]
}
