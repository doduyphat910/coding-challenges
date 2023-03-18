package seanprashad

func backspaceCompare(s string, t string) bool {
	return remove(s) == remove(t)
}

func remove(s string) string {
	var (
		i       int
		lengthS = len(s)
	)
	for i < lengthS {
		if string(s[i]) == "#" {
			if i == 0 {
				s = s[i+1:]
				lengthS -= 1
				continue
			}
			s = s[:i-1] + s[i+1:]
			i -= 2
			lengthS -= 2
		}
		i++
	}
	return s
}
