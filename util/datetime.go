package util

import (
	"time"

	"google.golang.org/genproto/googleapis/type/datetime"
)

func ToDateTime(t time.Time) *datetime.DateTime {

	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}

func FromDateTime(dt *datetime.DateTime) time.Time {
	return time.Date(
		int(dt.Year),
		time.Month(dt.Month),
		int(dt.Day),
		int(dt.Hours),
		int(dt.Minutes),
		int(dt.Seconds),
		int(dt.Nanos),
		time.UTC,
	)
}
