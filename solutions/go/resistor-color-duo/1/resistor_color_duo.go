package resistorcolorduo

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

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return ColorCode(colors[0])*10 + ColorCode(colors[1])
}
