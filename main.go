package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/yazeedalkhalaf/alyamamah-sp/handlers"
)

func main() {
	http.HandleFunc("/tweet-semester-progress", handlers.TweetSemesterProgress)
	http.HandleFunc("/latest-calendar", handlers.LatestCalendar)
	http.HandleFunc("/all-calendars", handlers.AllCalendars)
	http.HandleFunc("/calendar-by-name", handlers.CalendarByName)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("starting server to listen on port: %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
