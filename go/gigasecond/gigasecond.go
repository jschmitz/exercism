// Package gigasecond adds 1000000000 seconds to a time and returns the time.
package gigasecond

import (
	"time"
)

const (
	testVersion = 4   //exercism test version
	gigasecond  = 1e9 //int64 nanosecond
)

// AddGigasecond inputs a time, adds a gigasecond and returns resulting time.
func AddGigasecond(t time.Time) time.Time {
	d := time.Duration(gigasecond) * time.Second
	return t.Add(d)
}
