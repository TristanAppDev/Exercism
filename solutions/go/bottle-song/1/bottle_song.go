package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	nums := map[int]string{
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

	lines := make([]string, 0, takeDown*4)

	for i := 0; i < takeDown; i++ {
		b := startBottles - i
		s := "bottles"
		if b == 1 {
			s = "bottle"
		}
		line := fmt.Sprintf("%s green %s hanging on the wall,", nums[b], s)
		lines = append(lines, line, line)
		lines = append(lines, "And if one green bottle should accidentally fall,")

		if b == 1 {
			line = "There'll be no green bottles hanging on the wall."
			lines = append(lines, line)
		} else {
			if b == 2 {
				s = "bottle"
			}
			line = fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(nums[b-1]), s)
			lines = append(lines, line)
		}

		if b > 1 && i < takeDown-1 {
			lines = append(lines, "")
		}
	}

	return lines
}
