package common

import (
	"strconv"
	"time"
)

func CurrentTimeEpoch() string {
	now := time.Now()

	return strconv.FormatInt(now.Unix(), 10)
}

func ParseTime(date string) (string, error) {

	var itemDate time.Time
	var error error
	if len(date) == 0 {
		// no date passed, so return current date
		itemDate = time.Now()
		return itemDate.Format(time.DateOnly), nil
	}
	itemDate, error = time.Parse("2006-01-02", date)
	if error != nil {
		itemDate, error = time.Parse("20060102", date)
		if error != nil {
			return "", error
		}
	}
	return itemDate.Format(time.DateOnly), nil
}
