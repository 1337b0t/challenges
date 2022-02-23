package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	parseDate, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return parseDate.UTC()
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	parseDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return time.Now().After(parseDate)

}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	parseDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	hour := parseDate.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dt := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %s.",
		dt.Weekday(), dt.Month(), dt.Day(), dt.Year(), dt.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
