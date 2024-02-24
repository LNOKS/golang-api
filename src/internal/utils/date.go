package utils

import "time"

const DateFormat = "2006-01-02"

func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

func FormatDayMonth(t time.Time) string {
	return t.Format("January 2")
}
