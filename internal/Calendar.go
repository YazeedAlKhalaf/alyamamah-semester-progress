package internal

import (
	"encoding/json"
	"os"
	"time"
)

const (
	layout string = "Jan 2, 2006"
)

type Calendar struct {
	Events []Event `json:"events"`
}

func Ternary[T any](condition bool, If, Else T) T {
	if condition {
		return If
	}

	return Else
}

func (c Calendar) GetEventsWithDate(date time.Time) ([]Event, error) {
	var events []Event

	for _, event := range c.Events {
		startDate, err := time.Parse(layout, event.StartDate)
		if err != nil {
			return []Event{}, err
		}

		var endDate *time.Time
		if event.EndDate != "" {
			someDate, err := time.Parse(layout, event.EndDate)
			if err != nil {
				return []Event{}, err
			}

			endDate = &someDate
		}

		if endDate == nil {
			if date.Year() == startDate.Year() && date.YearDay() == startDate.YearDay() && date.Month() == date.Month() {
				events = append(events, event)
			}
		} else {
			if startDate.Before(date) && endDate.After(date) {
				events = append(events, event)
			}
		}

	}

	return events, nil
}

func NewCalendarFromFile(path string) (Calendar, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Calendar{}, nil
	}

	var calendar Calendar

	err = json.Unmarshal(data, &calendar)
	if err != nil {
		return Calendar{}, nil
	}

	return calendar, nil
}
