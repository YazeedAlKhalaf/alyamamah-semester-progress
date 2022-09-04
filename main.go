package main

import (
	"fmt"
	"os"
	"time"

	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func main() {
	calendar, err := internal.NewCalendarFromFile("alyamamah-calendar.json")
	if err != nil {
		fmt.Println("error creating calendar from file:", err)
		os.Exit(1)
	}
	//fmt.Println("calendar:", calendar)
	events, err := calendar.GetEventsWithDate(
		// time.Date(
		// 	2022,
		// 	8,
		// 	28,
		// 	0,
		// 	0,
		// 	0,
		// 	0,
		// 	&time.Location{},
		// ),
		time.Now(),
	)
	if err != nil {
		fmt.Println("error creating calendar from file:", err)
		os.Exit(1)
	}

	fmt.Println("events today:", events)
}
