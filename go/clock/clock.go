package clock

import "fmt"

// Clock represents time
type Clock struct {
	hh int
	mm int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hh, c.mm)
}

// Add should add time to the given clock
func (c Clock) Add(mins int) Clock {
	c.hh += mins / 60
	c.mm += mins % 60
	return c.normalize()
}

// Subtract should subtract time to the given clock
func (c Clock) Subtract(mins int) Clock {
	c.hh -= mins / 60
	c.mm -= mins % 60
	return c.normalize()
}

func (c Clock) normalize() Clock {
	hoursToSubtract := 0
	for c.mm < 0 {
		c.mm = 60 + c.mm
		hoursToSubtract++
	}
	remainingM := c.mm % 60
	c.hh += c.mm / 60
	c.hh -= hoursToSubtract
	for c.hh < 0 {
		c.hh = 24 + c.hh
	}
	if c.hh >= 24 {
		c.hh = c.hh % 24
	}
	c.mm = remainingM
	return c
}

// New constructs a new clock
func New(hour, mins int) Clock {
	c := Clock{hour, mins}
	return c.normalize()
}
