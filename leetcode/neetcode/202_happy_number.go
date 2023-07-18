package neetcode

func isHappy(n int) bool {
	var nMap = make(map[int]bool)
	_, ok := nMap[n]
	for !ok {
		nMap[n] = true
		n = sum(n)
		if n == 1 {
			return true
		}
		_, ok = nMap[n]
	}

	return false
}

func sum(n int) int {
	var sum int
	for n >= 10 {
		num := n % 10
		sum += num * num
		n /= 10
	}
	return n*n + sum
}
