package cmdutil

import "time"

const displayTimeFormatLayout = "2006-01-02 15:04:05"

func FormatTimeForDisplay(t time.Time) string {
	return t.Format(displayTimeFormatLayout)
}
