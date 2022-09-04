package internal

import "github.com/dghubble/go-twitter/twitter"

func CreateTweet(twitterClient *twitter.Client, tweetText string) (*twitter.Tweet, error) {
	tweet, _, err := twitterClient.Statuses.Update(tweetText, nil)
	if err != nil {
		return &twitter.Tweet{}, err
	}

	return tweet, nil
}
