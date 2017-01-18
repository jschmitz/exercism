// Package clock implements a Clock with continous roll over
package clock

import (
	"fmt"
)

const testVersion = 4

type Clock int

// New accepts hour and minute ints and
// converts to Clock minutes - (24 * 60) max
func New(hour, minute int) Clock {
	minutes := (hour*60 + minute)
	return minutesToClock(minutes)
}

// String formats the Clock to hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add updates Clock by minutes
func (c Clock) Add(minutes int) Clock {
	return minutesToClock(int(c) + minutes)
}

func minutesToClock(minutes int) Clock {
	minutes = minutes % (24 * 60)
	if minutes < 0 {
		minutes += (24 * 60)
	}
	return Clock(minutes)

}
