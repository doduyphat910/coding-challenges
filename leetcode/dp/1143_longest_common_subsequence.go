package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp = make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
				continue
			}
			dp[i][j] = dp[i+1][j]
			if dp[i+1][j] < dp[i][j+1] {
				dp[i][j] = dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}

// 2D dp
// O(n*m)
