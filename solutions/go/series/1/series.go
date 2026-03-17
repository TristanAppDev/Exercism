package series

func All(n int, s string) []string {
	result := make([]string, 0, len(s))
	for i := 0; n <= len(s); i++ {
		result = append(result, s[i:n])
		n++
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}
