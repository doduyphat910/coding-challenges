package dp

func wordBreak(s string, wordDict []string) bool {
	var dp = make([]bool, len(s)+2)
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for j := range wordDict {
			if i+len(wordDict[j]) <= len(s) &&
				wordDict[j] == string(s[i:i+len(wordDict[j])]) &&
				dp[i+len(wordDict[j])] {
				dp[i] = true
			}
		}
	}

	return dp[0]
}
