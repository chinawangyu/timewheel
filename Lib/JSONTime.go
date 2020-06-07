package Lib

import (
	"strings"
	"time"
)

// JSONTime format json time field by myself

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

var ZERO_TIME Time

func (t *Time) UnmarshalJSON(dateTimeByte []byte) (err error) {
	dateTimeStr := strings.Trim(string(dateTimeByte), `"`)
	if "0000-00-00 00:00:00" == dateTimeStr {
		*t = ZERO_TIME
		return nil
	}
	now, err := time.ParseInLocation(timeFormart, dateTimeStr, time.Local)
	*t = Time(now)
	return
}
