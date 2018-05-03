// https://golang.org/pkg/time/#Now
package lib

import "time"

// Now returns the a timestamp.
func Now() string {
	return time.Now().Format("20060102150405")
}

// NowUTC returns the a UTC formatted timestamp.
func NowUTC() string {
	return nowUTC := time.Now().UTC().Format("20060102150405")
}
