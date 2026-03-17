package triangle

type Kind string

const (
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
	NaT Kind = "NaT"
)

// KindFromSides slightly faster version than my first ones
func KindFromSides(a, b, c float64) Kind {
	// a and b will be smaller than c
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}

	switch {
	case a <= 0: // negative side length not allowed
		return NaT
	case a+b < c: // a+b have to be greater than c
		return NaT
	case a == c && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
