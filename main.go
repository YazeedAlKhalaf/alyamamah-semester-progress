package main

import (
	"net/http"

	"github.com/yazeedalkhalaf/alyamamah-sp/cloudfunctions"
)

func main() {
	http.HandleFunc("/tweet-semester-progress", cloudfunctions.TweetSemesterProgress)

	http.ListenAndServe(":8081", nil)
}
