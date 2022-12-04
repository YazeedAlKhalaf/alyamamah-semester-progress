package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func TweetSemesterProgress(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		fmt.Println("someone tried to access the API without a token.")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	strings.Split(authHeader, "Bearer ")
	isAuthorized, err := internal.IsAuthorized(strings.Split(authHeader, "Bearer ")[1])
	if err != nil {
		fmt.Println("something went wrong running isAuthorized:", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if !isAuthorized {
		fmt.Println("someone tried to access the API with a token maliciously:", authHeader)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	twitterClient, err := internal.CreateTwitterClient()
	if err != nil {
		fmt.Println("something went wrong creating the twitter client:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	calendar, err := internal.NewCalendarFromLatestFile()
	if err != nil {
		fmt.Println("something went wrong reading latest calendar from file:", err)
		w.WriteHeader(http.StatusInternalServerError)
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

	currentWeek := math.Ceil(float64(currentDay) / 7)

	progressBar := internal.CreateProgressBar(float32(currentDay), float32(totalDays), 15)

	tweetText := fmt.Sprintln(
		progressBar,
		fmt.Sprintln(""),
		fmt.Sprintln(""),
		fmt.Sprintf("🗓️ %d/%d days passed", currentDay, totalDays),
		fmt.Sprintln(""),
		fmt.Sprintf("Current Week %d", int(currentWeek)),
		fmt.Sprintln(""),
		fmt.Sprintln(""),
		`#alyamamah`,
	)

	var tweets []*twitter.Tweet
	tweet, err := internal.CreateTweet(twitterClient, tweetText)
	if err != nil {
		fmt.Println("something went wrong creating the tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tweets = append(tweets, tweet)
	fmt.Println("tweet has been published:", tweet.Text)

	// publish events as replies to tweets
	todayEvents, err := calendar.GetEventsWithDate(time.Now())
	if err != nil {
		fmt.Println("something went wrong getting today events:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for index, event := range todayEvents {
		startAndEndDays := strings.Split(event.Day, "-")
		startDay := startAndEndDays[0]
		startDate := event.StartDate
		doesContainEndDay := len(startAndEndDays) == 2
		endDay := internal.Ternary(doesContainEndDay, startAndEndDays[1], startDay)
		endDate := internal.Ternary(doesContainEndDay, event.EndDate, startDate)
		eventText := fmt.Sprintln(
			fmt.Sprintf("🏛️ Event %d:", index+1),
			fmt.Sprintln(""),
			fmt.Sprintln(""),
			"• ",
			event.Name,
			fmt.Sprintln(""),
			fmt.Sprintln(""),
			fmt.Sprintf("Start Date: %s - %s", startDay, startDate),
			fmt.Sprintf("End Date: %s - %s", endDay, endDate),
		)

		replyTweet, err := internal.ReplyToTweet(twitterClient, eventText, tweets[len(tweets)-1].ID)
		if err != nil {
			fmt.Println("something went wrong replying to the tweet:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tweets = append(tweets, replyTweet)
		fmt.Println("reply tweet has been published:", replyTweet.Text)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{
		"tweets": tweets,
	})
	if err != nil {
		fmt.Println("something went wrong encoding the tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
