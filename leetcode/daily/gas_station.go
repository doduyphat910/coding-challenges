package daily

func canCompleteCircuit(gas []int, cost []int) int {
	var sum, result, totalCost int
	for i := range gas {
		sum += gas[i] - cost[i]
		totalCost += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			result = i + 1
		}
	}

	if result == len(gas) || totalCost < 0 {
		return -1
	}

	return result
}
