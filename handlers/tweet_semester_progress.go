package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func TweetSemesterProgress(w http.ResponseWriter, r *http.Request) {
	twitterClient, err := internal.CreateTwitterClient()
	if err != nil {
		fmt.Println("something went wrong creating the twitter client:", err)
		return
	}

	calendar, err := internal.NewCalendarFromFile("alyamamah-calendar.json")
	if err != nil {
		fmt.Println("something went wrong reading calendar from file:", err)
		return
	}
	currentDay, err := calendar.GetCurrentDayInSemester()
	if err != nil {
		fmt.Println("something went wrong getting the current day in semester:", err)
		return
	}
	totalDays, err := calendar.GetTotalDaysNumber()
	if err != nil {
		fmt.Println("something went wrong getting total days number:", err)
		return
	}

	progressBar := internal.CreateProgressBar(float32(currentDay), float32(totalDays), 15)

	tweetText := fmt.Sprintln(
		progressBar,
		fmt.Sprintln(""),
		fmt.Sprintln(""),
		fmt.Sprintf("🗓️ %d/%d days passed", currentDay, totalDays),
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
