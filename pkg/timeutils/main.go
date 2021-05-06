package timeutils

import (
	"fmt"
	"time"
)

func ClosestMonday(date string) (time.Time, error) {
	layout := "2006-01-02"
	currentDate, err := time.Parse(layout, date)

	if err != nil {
		return time.Time{}, fmt.Errorf("Not a valid date: %q", date)
	}

	for {
		if currentDate.Weekday().String() == "Monday" {
			break
		}

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return currentDate, nil
}
