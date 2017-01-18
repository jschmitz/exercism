// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
	"fmt"
)

const (
	testVersion   = 4  //exercism test version
	minutesInHour = 60 // number of minutes in a hour
	hoursInDay    = 24 // number of hours in a day
)

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	remainderMinutes, minuteHours := continousMinutes(minute)
	h := continousHours(minuteHours + hour)
	return Clock{hour: h, minute: remainderMinutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	remainderMinutes, minuteHours := continousMinutes(c.minute + minutes)
	h := continousHours(minuteHours + c.hour)
	c.minute = remainderMinutes
	c.hour = h
	return c
}

func continousHours(hours int) int {
	if hours < 0 {
		hours = hoursInDay - (hours*-1)%hoursInDay
		if hours == hoursInDay {
			hours = 0
		}
	} else {
		hours = hours % hoursInDay
	}

	return hours
}

func continousMinutes(minutes int) (int, int) {
	var remainder, hours int
	if minutes < 0 {
		remainder = minutesInHour - (minutes*-1)%minutesInHour
		hours = (minutes / minutesInHour) - 1
	} else {
		remainder = minutes % minutesInHour
		hours = minutes / minutesInHour
	}

	return remainder, hours
}
