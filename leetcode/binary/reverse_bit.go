package binary

func reverseBits(num uint32) uint32 {
	var (
		result uint32
		j      = uint32(1 << 31)
	)

	for num != 0 {
		if (num & 1) != 0 {
			result = result | j
		}

		j = j >> 1
		num = num >> 1

	}

	return result
}
