package main

import (
	"fmt"
	"time"
)

type Clock struct {
	Time int
}

func New(hour, minute int) Clock {
	m_quotient := minute / 60
	m_remainder := minute % 60
	h_remainder := (hour + m_quotient) % 24

	if m_remainder < 0 {
		m_remainder = 60 + m_remainder
		h_remainder -= 1
	}

	if h_remainder < 0 {
		h_remainder = 24 + h_remainder
	}

	return Clock{Time: h_remainder + m_remainder}
}

func (c Clock) String() string {
	sec := int64(c.Time) / 1000
	msec := int64(c.Time) % 1000
	return time.Unix(sec, msec*int64(time.Millisecond)).Format("15:04")
}

func (c Clock) Add(minutes int) Clock {
	sec := int64(c.Time) / 1000
	msec := int64(c.Time) % 1000
	newTime := time.Unix(sec, msec*int64(time.Millisecond)).Add(time.Duration(minutes))

	return New(newTime.Hour(), newTime.Minute())
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func main() {
	c := New(201, 3001)
	fmt.Println(c.String())

}
