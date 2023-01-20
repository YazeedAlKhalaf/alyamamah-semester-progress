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
	Title     string  `json:"title"`
	Events    []Event `json:"events"`
	IsCurrent bool    `json:"isCurrent"`
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

	dateDiff := int(lastDate.Sub(firstDate).Hours() / 24)
	dateDiff = int(math.Abs(float64(dateDiff)))
	return dateDiff, nil
}

func (c Calendar) GetCurrentDayInSemester() (int, error) {
	firstDate, err := c.GetFirstDay()
	if err != nil {
		return -1, err
	}

	semesterDays, err := c.GetTotalDaysNumber()
	if err != nil {
		return -1, err
	}

	doneDuration := time.Now().Sub(firstDate)
	doneDays := int(math.Ceil(doneDuration.Hours() / 24))

	// If the done days is negative, it means that the semester has not started yet.
	// --- now --- firstDate --- lastDate ---
	// ---  13  ---    15    ---    25    ---
	// If this drawing increases from left to right, then the semester has not started yet.
	// If now = 13 and firstDate = 15, then doneDays = -2, but since the math.Ceil function rounds up, it will be -1
	// even if the doneDays is -1.5, that is because it always goes up.
	// But in our case, want to know the remaining days in negative numbers, so we subtract 1 from the doneDays.
	// We can use math.Floor with a check on doneDuration, but I think this is more readable.
	if doneDays < 0 {
		doneDays -= 1
	}

	// If the done days is greater than the semester days, it means that the semester has ended.
	// --- firstDate --- lastDate --- now ---
	// ---     15    ---    25    ---  29 ---
	// If this drawing increases from left to right, then the semester has ended.
	// If now = 29 and firstDate = 15, then doneDays = 14, we check if the doneDays is more than the semesterDays.
	// Since semesterDays in this example is lastDate - firstDate = 10, we set the doneDays to 10.
	if doneDays > semesterDays {
		doneDays = semesterDays
	}

	return doneDays, nil
}

func (c Calendar) IsCurrentDayInSemester() (bool, error) {
	firstDate, err := c.GetFirstDay()
	if err != nil {
		return false, err
	}

	lastDate, err := c.GetLastDay()
	if err != nil {
		return false, err
	}

	now := time.Now()
	if now.After(firstDate) && now.Before(lastDate) {
		return true, nil
	}

	return false, nil
}

func newCalendarFromFile(path string) (Calendar, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Calendar{}, err
	}

	var calendar Calendar

	err = json.Unmarshal(data, &calendar)
	if err != nil {
		return Calendar{}, err
	}

	isCurrent, err := calendar.IsCurrentDayInSemester()
	if err != nil {
		return Calendar{}, err
	}

	calendar.IsCurrent = isCurrent

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
