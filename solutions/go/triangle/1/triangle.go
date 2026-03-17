package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if !isValid(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

func isValid(a, b, c float64) bool {
	r := (a + b + c) * (-a + b + c) * (a - b + c) * (a + b - c) // derived from Heron's formula
	return r > 0
}
