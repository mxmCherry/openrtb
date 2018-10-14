package openrtb3

import (
	"fmt"
	"time"
)

// timeFormats holds possible ISO-8601 date/time formats (https://en.wikipedia.org/wiki/ISO_8601).
var timeFormats = []string{
	time.RFC3339,
	"2006-01-02",
	"2006-01-02T15:04:05+00:00",
	"2006-01-02T15:04:05Z",
	"20060102T150405Z",
}

// ParseTime parses a date/time string.
func ParseTime(s string) (time.Time, error) {
	var err error
	for _, f := range timeFormats {
		if t, e := time.Parse(f, s); e == nil {
			return t, nil
		} else {
			err = e
		}
	}
	return time.Time{}, fmt.Errorf("openrtb3: ParseTime(%q) failed: %s", s, err)
}
