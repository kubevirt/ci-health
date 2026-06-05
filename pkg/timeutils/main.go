package timeutils

import (
	"fmt"
	"time"
)

func ClosestMonday(date string) (time.Time, error) {
	layout := "2006-01-02"
	currentDate, err := time.Parse(layout, date)

	if err != nil {
		return time.Time{}, fmt.Errorf("not a valid date: %q", date)
	}

	for currentDate.Weekday() != time.Monday {
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return currentDate, nil
}
