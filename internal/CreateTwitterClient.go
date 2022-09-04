package internal

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// CreateTwitterClient uses the environment variables to create a twitter authenticated client.
func CreateTwitterClient() (*twitter.Client, error) {
	envConfig, err := LoadConfig(".")
	if err != nil {
		return &twitter.Client{}, err
	}

	config := oauth1.NewConfig(envConfig.ConsumerKey, envConfig.ConsumerKeySecret)
	token := oauth1.NewToken(envConfig.AccessToken, envConfig.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	twitterClient := twitter.NewClient(httpClient)
	return twitterClient, nil
}
