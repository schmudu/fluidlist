package common

import (
	"strconv"
	"time"
)

func CurrentTimeEpoch() string {
	now := time.Now()

	return strconv.FormatInt(now.Unix(), 10)
}

func CurrentTime() string {
	now := time.Now()

	// Output: "2026-11-09" (based on current date)
	return now.Format(time.DateOnly)
}
