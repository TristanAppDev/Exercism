package resistorcolortrio

import (
	"math"
	"strconv"
)

const (
	kilo int = 1_000
	mega int = 1_000_000
	giga int = 1_000_000_000
)

var colorList = []string{
	"black", "brown", "red", "orange", "yellow",
	"green", "blue", "violet", "grey", "white",
}

var colorMap = func() map[string]uint {
	m := make(map[string]uint, len(colorList))
	for i, color := range colorList {
		m[color] = uint(i)
	}
	return m
}()

func ColorCode(color string) int {
	return int(colorMap[color])
}

func Value(colors []string) int {
	return ColorCode(colors[0])*10 + ColorCode(colors[1])
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	number := Value(colors[:2])
	multiplier := int(math.Pow(10, float64(ColorCode(colors[2]))))
	ohms := number * multiplier

	switch {
	case ohms >= giga:
		return strconv.Itoa(ohms/giga) + " gigaohms"
	case ohms >= mega:
		return strconv.Itoa(ohms/mega) + " megaohms"
	case ohms >= kilo:
		return strconv.Itoa(ohms/kilo) + " kiloohms"
	default:
		return strconv.Itoa(ohms) + " ohms"
	}
}
