package util

import "time"

func DateTime(dateString string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	return timeTime(dateString, layoutFormat)
}

func Date(dateString string) time.Time {
	layoutFormat := "2006-01-02"
	return timeTime(dateString, layoutFormat)
}

func timeTime(layoutFormat, dateString string) time.Time {
	date, _ := time.Parse(layoutFormat, dateString)
	return date
}
