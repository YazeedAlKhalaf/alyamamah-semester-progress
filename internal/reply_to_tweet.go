package internal

import "github.com/dghubble/go-twitter/twitter"

func ReplyToTweet(twitterClient *twitter.Client, tweetText string, tweetID int64) (*twitter.Tweet, error) {
	tweet, _, err := twitterClient.Statuses.Update(tweetText, &twitter.StatusUpdateParams{
		InReplyToStatusID: tweetID,
	})
	if err != nil {
		return &twitter.Tweet{}, err
	}

	return tweet, nil
}
