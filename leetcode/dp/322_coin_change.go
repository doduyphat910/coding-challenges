package dp

func coinChange(coins []int, amount int) int {
	var dp = make(map[int]int)
	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			_, ok := dp[i]
			_, ok2 := dp[i-coins[j]]
			if ok2 && (!ok || 1+dp[i-coins[j]] < dp[i]) {
				dp[i] = 1 + dp[i-coins[j]]
			}
		}
	}

	c, ok := dp[amount]
	if !ok {
		return -1
	}
	return c
}

// time: O(a*c)
// space: O(a)
