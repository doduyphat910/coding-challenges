package other

func isMatch(s string, p string) bool {
	var dp = make(map[[2]int]bool)
	return dfs(0, 0, s, p, dp)
}

func dfs(i, j int, s, p string, dp map[[2]int]bool) bool {
	if _, ok := dp[[2]int{i, j}]; ok {
		return dp[[2]int{i, j}]
	}
	if i >= len(s) && j >= len(p) {
		return true
	}
	if j >= len(p) {
		return false
	}

	var match = i < len(s) && (s[i] == p[j] || string(p[j]) == ".")
	if j+1 < len(p) && string(p[j+1]) == "*" {
		dp[[2]int{i, j}] = (match && dfs(i+1, j, s, p, dp)) || dfs(i, j+2, s, p, dp)
		return dp[[2]int{i, j}]
	}
	if match {
		dp[[2]int{i, j}] = dfs(i+1, j+1, s, p, dp)
		return dp[[2]int{i, j}]
	}

	return false
}
