package binary

func getSum(a int, b int) int {
	var c int
	for b != 0 {
		c = a & b
		a = a ^ b
		b = (c) << 1
	}

	return a
}
