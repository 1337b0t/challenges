package clock

import (
	"fmt"
)

type Clock struct {
	Minutes int
}

func New(hour, minute int) Clock {
	dayMinutes := 24 * 60
	totalMinutes := (hour*60 + minute) % dayMinutes

	if totalMinutes < 0 {
		totalMinutes += dayMinutes
	}

	return Clock{Minutes: totalMinutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Minutes/60, c.Minutes%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.Minutes+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
