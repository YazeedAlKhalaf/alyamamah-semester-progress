package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func LatestCalendar(w http.ResponseWriter, r *http.Request) {
	calendar, err := internal.NewCalendarFromLatestFile()
	if err != nil {
		fmt.Println("something went wrong reading latest calendar from file:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"calendar": calendar,
	})
	if err != nil {
		fmt.Println("something went wrong encoding the tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
