package common

import (
	"errors"
	"strconv"
	"time"
)

func CurrentTimeEpoch() string {
	now := time.Now()

	return strconv.FormatInt(now.Unix(), 10)
}

func ParseTime(date string) (time.Time, bool, error) {
	// date entered will be either
	// 1. In the future, which will be return an error
	// 2. Empty, which will return time.now
	// 3. In the past, which will require a binary insertion
	// return parameters
	// string - date in string format
	// bool - whether or not item needs to be inserted in the past
	// error - error with the input

	var itemDate time.Time
	var nowDate = time.Now()
	var error error
	if len(date) == 0 {
		// no date passed, so return current date
		return nowDate, false, nil
	}
	itemDate, error = time.Parse("2006-01-02", date)
	if error != nil {
		itemDate, error = time.Parse("20060102", date)
		if error != nil {
			itemDate, error = time.Parse("2006/01/02", date)
			if error != nil {
				return nowDate, false, error
			}
		}
	}

	// validate that date provided is not after today
	if itemDate.After(nowDate) {
		return nowDate, false, errors.New("Item date is in the future.")
	}
	return itemDate, true, nil
}
