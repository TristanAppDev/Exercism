package atbash

import "strings"

var cipher = map[string]string{
	"a": "z",
	"b": "y",
	"c": "x",
	"d": "w",
	"e": "v",
	"f": "u",
	"g": "t",
	"h": "s",
	"i": "r",
	"j": "q",
	"k": "p",
	"l": "o",
	"m": "n",
	"n": "m",
	"o": "l",
	"p": "k",
	"q": "j",
	"r": "i",
	"s": "h",
	"t": "g",
	"u": "f",
	"v": "e",
	"w": "d",
	"x": "c",
	"y": "b",
	"z": "a",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
}

func Atbash(s string) string {
	var result string
	var counter int
	s = strings.ToLower(s)
	s = strings.Trim(s, ".")

	for i, r := range s {
		if r == ' ' {
			continue
		}
		if r >= 'a' && r <= 'z' || r >= '1' && r <= '9' {
			result += cipher[string(r)]
			counter++
			if counter%5 == 0 && i < len(s)-1 {
				result += " "
			}
		}
	}

	return result
}
