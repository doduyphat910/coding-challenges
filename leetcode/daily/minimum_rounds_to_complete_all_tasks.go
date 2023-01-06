package daily

import "sort"

func minimumRounds(tasks []int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})
	var (
		count        int
		taskByNumMap = make(map[int]int)
	)
	taskByNumMap[tasks[0]] = 1
	for i := 1; i < len(tasks); i++ {
		num, ok := taskByNumMap[tasks[i]]
		if !ok {
			num = 0
		}
		num += 1
		taskByNumMap[tasks[i]] = num
		if tasks[i] == tasks[i-1] {
			continue
		}
		if taskByNumMap[tasks[i-1]] == 1 {
			return -1
		}
		if taskByNumMap[tasks[i-1]]%3 == 0 {
			count += taskByNumMap[tasks[i-1]] / 3
		} else {
			count += taskByNumMap[tasks[i-1]]/3 + 1
		}
	}

	if taskByNumMap[tasks[len(tasks)-1]] == 1 {
		return -1
	}
	if taskByNumMap[tasks[len(tasks)-1]]%3 == 0 {
		count += taskByNumMap[tasks[len(tasks)-1]] / 3
	} else {
		count += taskByNumMap[tasks[len(tasks)-1]]/3 + 1
	}

	return count
}
