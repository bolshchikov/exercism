package clock

import "fmt"

const minutesInHour = 60
const minutesInDay = 24 * 60

// Clock object
type Clock struct {
	mins int
}

func (c Clock) String() string {
	hours := c.mins / minutesInHour
	minutes := c.mins % minutesInHour
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add should add time to the given clock
func (c Clock) Add(mins int) Clock {
	c.mins += mins
	return c.normalize()
}

// Subtract should subtract time to the given clock
func (c Clock) Subtract(mins int) Clock {
	c.mins -= mins
	return c.normalize()
}

func (c Clock) normalize() Clock {
	for c.mins < 0 || c.mins >= minutesInDay {
		switch {
		case c.mins < -minutesInDay || c.mins > minutesInDay:
			c.mins = c.mins % minutesInDay
		case -minutesInDay < c.mins && c.mins < 0:
			c.mins = minutesInDay + c.mins
		case c.mins == minutesInDay:
			c.mins = 0
		}
	}
	return c
}

// New creates new clock object
func New(hh, mm int) Clock {
	c := Clock{hh*minutesInHour + mm}
	return c.normalize()
}
