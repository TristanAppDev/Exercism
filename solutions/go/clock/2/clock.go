package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock int

var minutesPerDay = 1440

func New(h, m int) Clock {
	return ToClock(h*60 + m)
}

func (c Clock) Add(m int) Clock {
	return ToClock(int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
	return ToClock(int(c) - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func ToClock(minutes int) Clock {
	if minutes < 0 {
		return Clock(minutesPerDay + (minutes % minutesPerDay))
	}
	return Clock(minutes % minutesPerDay)
}
