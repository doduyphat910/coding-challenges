package string

func isValid(s string) bool {
	defaultMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := make([]string, 0)

	for i := range s {
		value, ok := defaultMap[string(s[i])]
		if ok {
			stack = append(stack, value)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != string(s[i]) {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
