package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func AllCalendars(w http.ResponseWriter, r *http.Request) {
	calendars, err := internal.GetAllCalendars()
	if err != nil {
		fmt.Println("something went wrong getting all calendars:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"calendars": calendars,
	})
	if err != nil {
		fmt.Println("something went wrong encoding the calendars:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
