package say

// translate number to words up to 1_000_000_000_000
var small = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var big = []string{
	"thousand",
	"million",
	"billion",
	"trillion",
}

func Say(n int64) (string, bool) {

	if n < 0 || n >= 1_000_000_000_000 {
		return "", false
	}

	if n < 20 {
		return small[n], true
	}

	return _Say(n), true
}

func _Say(n int64) string {
	switch {
	case n < 20:
		return small[n]
	case n < 100:
		word := tens[n/10-2]
		rem := n % 10
		if rem > 0 {
			word += "-" + small[rem]
		}
		return word
	case n < 1000:
		word := small[n/100] + " hundred"
		rem := n % 100
		if rem > 0 {
			word += " " + _Say(rem)
		}
		return word
	}

	word := ""

	if rem := n % 1000; rem > 0 {
		word = _Say(rem)
	}
	for i := 0; n >= 1000; i++ {
		n /= 1000
		if rem := n % 1000; rem > 0 {
			ix := _Say(rem) + " " + big[i]
			if word > "" {
				ix += " " + word
			}
			word = ix
		}
	}

	return word
}
