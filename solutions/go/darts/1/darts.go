package darts

import "math"

func Score(x, y float64) int {
	dart_pos := math.Sqrt(x*x + y*y)

	switch {
	case dart_pos <= 1.0:
		return 10
	case dart_pos <= 5.0:
		return 5
	case dart_pos <= 10.0:
		return 1
	default:
		return 0
	}
}
