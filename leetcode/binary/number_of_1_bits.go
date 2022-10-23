package binary

func hammingWeight(num uint32) int {
	total := 0
	for num != 0 {
		lastBit := num & 1
		if lastBit != 0 {
			total += 1
		}
		num = num >> 1
	}

	return total
}
