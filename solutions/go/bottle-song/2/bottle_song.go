package bottlesong

import (
	"fmt"
	"strings"
)

var nums = map[int]string{
	10: "Ten",
	9:  "Nine",
	8:  "Eight",
	7:  "Seven",
	6:  "Six",
	5:  "Five",
	4:  "Four",
	3:  "Three",
	2:  "Two",
	1:  "One",
}

func Recite(startBottles, takeDown int) []string {
	lines := []string{}
	for i := startBottles; i > startBottles-takeDown; i -= 1 {
		lines = append(lines, verse(i)...)

		if i > startBottles-takeDown+1 {
			lines = append(lines, "")
		}
	}
	return lines
}

func verse(n int) []string {
	switch n {
	case 1:
		return []string{
			"One green bottle hanging on the wall,",
			"One green bottle hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be no green bottles hanging on the wall.",
		}
	case 2:
		return []string{
			"Two green bottles hanging on the wall,",
			"Two green bottles hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be one green bottle hanging on the wall.",
		}
	default:
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", nums[n]),
			fmt.Sprintf("%s green bottles hanging on the wall,", nums[n]),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(nums[n-1])),
		}
	}
}
