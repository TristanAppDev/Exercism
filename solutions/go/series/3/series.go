package series

func All(n int, s string) []string {
	runes := []rune(s)
	rLen := len(runes)
	res := make([]string, 0, rLen-n+1)

	for i := 0; n <= rLen; i++ {
		res = append(res, string(runes[i:n]))
		n++
	}
	return res
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
