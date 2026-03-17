// Package gigasecond should have a package comment that summarizes what it's about.
package gigasecond

// import path for the time package from the standard library
import "time"

const gigaseconds = 1_000_000_000

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigaseconds)
}
