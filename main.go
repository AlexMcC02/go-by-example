package main

import (
	"fmt"
	"time"
)

// Go offers extensive support for times and durations.

func main() {
	p := fmt.Println

	// Getting the current time.
	now := time.Now()
	p(now)

	// You can build a time struct by providing the year, month, day, etc.
	// Times are ALWAYS associated with a timezone, here UTC is provided.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651386237, time.UTC)
	p(then)

	// You can extract the various componets of the time value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Monday-Sunday is also available.
	p(then.Weekday())

	// These methods will compare two times.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The sub method returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use Add to advance time.
	p(then.Add(diff))
	p(then.Add(-diff))
}