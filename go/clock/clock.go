package clock

import "fmt"

const minsInHour = 60
const hoursInDay = 24

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
	c.hh += mins / minsInHour
	c.mm += mins % minsInHour
	return c.normalize()
}

// Subtract should subtract time to the given clock
func (c Clock) Subtract(mins int) Clock {
	c.hh -= mins / minsInHour
	c.mm -= mins % minsInHour
	return c.normalize()
}

func (c Clock) normalize() Clock {
	hoursToSubtract := 0
	for c.mm < 0 {
		c.mm = minsInHour + c.mm
		hoursToSubtract++
	}
	c.hh = c.hh + c.mm/minsInHour - hoursToSubtract
	for c.hh < 0 {
		c.hh = hoursInDay + c.hh
	}
	if c.hh >= hoursInDay {
		c.hh = c.hh % hoursInDay
	}
	c.mm = c.mm % minsInHour
	return c
}

// New constructs a new clock
func New(hour, mins int) Clock {
	c := Clock{hour, mins}
	return c.normalize()
}
