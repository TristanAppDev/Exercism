package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	h, m int
}

func New(h, m int) Clock {

	// normalize negative hours
	if h < 0 {
		h = toPosHours(h)
	}

	if m < 0 {
		mH := m/60 - 1
		m = (m % 60) + 60
		h = toPosHours(h + mH)
	}

	h = ((h % 24) + (m / 60)) % 24
	m = (m % 60)

	return Clock{h: h, m: m}
}

func toPosHours(h int) int {
	return (h % 24) + 24
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
