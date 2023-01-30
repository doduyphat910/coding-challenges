package daily

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}

	var (
		peopleByTrustedMap = make(map[int][]int)
		trustMap           = make(map[int][]int)
	)
	for i := range trust {
		people, ok := peopleByTrustedMap[trust[i][1]]
		if !ok {
			people = make([]int, 0)
		}
		people = append(people, trust[i][0])
		peopleByTrustedMap[trust[i][1]] = people

		trusts, ok := trustMap[trust[i][0]]
		if !ok {
			trusts = make([]int, 0)
		}
		trusts = append(trusts, trust[i][1])
		trustMap[trust[i][0]] = trusts
	}

	for i := range peopleByTrustedMap {
		_, ok := trustMap[i]
		if len(peopleByTrustedMap[i]) == n-1 && !ok {
			return i
		}
	}

	return -1
}
