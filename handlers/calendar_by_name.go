package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func CalendarByName(w http.ResponseWriter, r *http.Request) {
	calendarTitle := r.URL.Query().Get("name")
	if calendarTitle == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calendars, err := internal.GetAllCalendars()
	if err != nil {
		fmt.Println("something went wrong getting all calendars:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	calendar := internal.FirstWhere(calendars, func(c *internal.Calendar) bool { return c.Title == calendarTitle })
	if calendar == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	currentDay, err := calendar.GetCurrentDayInSemester()
	if err != nil {
		fmt.Println("something went wrong getting the current day in semester:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalDays, err := calendar.GetTotalDaysNumber()
	if err != nil {
		fmt.Println("something went wrong getting total days number:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"currentDay": currentDay,
		"totalDays":  totalDays,
		"calendar":   calendar,
	})
	if err != nil {
		fmt.Println("something went wrong encoding the calendars:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
