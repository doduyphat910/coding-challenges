package neetcode

func plusOne(digits []int) []int {
	var remember int
	digits[len(digits)-1] += 1
	if digits[len(digits)-1] >= 10 {
		digits[len(digits)-1] %= 10
		remember = 1
	}

	for i := len(digits) - 2; i >= 0; i-- {
		digits[i] += remember
		remember = 0
		if digits[i] >= 10 {
			remember = 1
			digits[i] %= 10
		}
	}

	if remember > 0 {
		digits = append([]int{remember}, digits...)
		return digits
	}

	return digits
}
