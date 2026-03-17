package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var cipher = make([]rune, len(plain))
	for i, r := range plain {
		if r >= 'a' && r <= 'z' {
			cipher[i] = rune('a' + (int(r-'a')+shiftKey)%26)
		} else if r >= 'A' && r <= 'Z' {
			cipher[i] = rune('A' + (int(r-'A')+shiftKey)%26)
		} else {
			cipher[i] = r
		}
	}
	return string(cipher)
}
