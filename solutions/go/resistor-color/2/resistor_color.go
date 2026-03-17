package resistorcolor

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

// Colors returns the list of all colors.
func Colors() []string {
	return colorList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return int(colorMap[color])
}
