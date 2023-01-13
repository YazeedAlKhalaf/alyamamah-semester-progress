package internal

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	layout       string = "Jan 2, 2006"
	calendarsDir string = "./calendars"
)

type Calendar struct {
	Title  string  `json:"title"`
	Events []Event `json:"events"`
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

func (c Calendar) GetFirstDay() (time.Time, error) {
	if len(c.Events) < 0 {
		return time.Time{}, fmt.Errorf("no events in the calendar")
	}

	chosenDate := c.Events[1].StartDate

	firstDate, err := time.Parse(layout, chosenDate)
	if err != nil {
		return time.Time{}, err
	}

	return firstDate, nil
}

func (c Calendar) GetLastDay() (time.Time, error) {
	lastDate, err := time.Parse(layout, c.Events[len(c.Events)-1].StartDate)
	if err != nil {
		return time.Time{}, err
	}

	return lastDate, nil
}

func (c Calendar) GetTotalDaysNumber() (int, error) {
	firstDate, err := c.GetFirstDay()
	if err != nil {
		return -1, err
	}

	lastDate, err := c.GetLastDay()
	if err != nil {
		return -1, err
	}

	dateDiff := int(lastDate.Sub(firstDate) / (24 * time.Hour))
	dateDiff = int(math.Abs(float64(dateDiff)))
	return dateDiff, nil
}

func (c Calendar) GetCurrentDayInSemester() (int, error) {
	lastDate, err := c.GetLastDay()
	if err != nil {
		return -1, err
	}

	semesterDays, err := c.GetTotalDaysNumber()
	if err != nil {
		return -1, err
	}

	nowLastDiff := int(lastDate.Sub(time.Now()) / (24 * time.Hour))
	nowLastDiff = int(math.Abs(float64(nowLastDiff)))

	currentDay := semesterDays - nowLastDiff

	firstDate, err := c.GetFirstDay()
	if err != nil {
		return -1, err
	}

	if firstDate.After(time.Now()) {
		currentDay *= -1
	}

	return currentDay, nil
}

func newCalendarFromFile(path string) (Calendar, error) {
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

func NewCalendarFromLatestFile() (Calendar, error) {
	calendarFiles, err := ioutil.ReadDir(calendarsDir)
	if err != nil {
		return Calendar{}, err
	}

	// This regular expression get the digits at the start of the line.
	// For more info visit https://regex101.com and text the regular expression.
	numRegExp := regexp.MustCompile(`^\d+`)
	maxNumber := -1
	calendarsMap := map[int]fs.FileInfo{}
	for _, calendarFile := range calendarFiles {
		if calendarFile.IsDir() {
			continue
		}

		numString := string(numRegExp.Find([]byte(calendarFile.Name())))
		num, err := strconv.Atoi(numString)
		if err != nil {
			continue
		}

		calendarsMap[num] = calendarFile

		if num > maxNumber {
			maxNumber = num
		}
	}
	if maxNumber == -1 {
		return Calendar{}, fmt.Errorf("no calendars, please check the directory you provided.")
	}

	return newCalendarFromFile(fmt.Sprintf("%s/%s", calendarsDir, calendarsMap[maxNumber].Name()))
}

func GetAllCalendars() ([]Calendar, error) {
	calendarFiles, err := ioutil.ReadDir(calendarsDir)
	if err != nil {
		return []Calendar{}, err
	}

	calendars := []Calendar{}
	for _, calendarFile := range calendarFiles {
		if calendarFile.IsDir() {
			continue
		}
		calendar, err := newCalendarFromFile(fmt.Sprintf("%s/%s", calendarsDir, calendarFile.Name()))
		if err != nil {
			continue
		}
		calendars = append(calendars, calendar)
	}

	return calendars, nil
}
