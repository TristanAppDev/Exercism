package series

func All(n int, s string) []string {
	rLen := len([]rune(s))
	res := make([]string, 0, rLen-n+1)

	for i := 0; n <= rLen; i++ {
		res = append(res, string(s[i:n]))
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
