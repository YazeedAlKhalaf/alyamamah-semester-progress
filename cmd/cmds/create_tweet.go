package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func init() {
	createTweetCmd.Flags().String("text", "", "Provides text that the tweet will contain.")
	rootCmd.AddCommand(createTweetCmd)
}

var createTweetCmd = &cobra.Command{
	Use:   "create-tweet",
	Short: "Create a tweet",
	Run: func(cmd *cobra.Command, args []string) {
		twitterClient, err := internal.CreateTwitterClient()
		if err != nil {
			fmt.Println(fmt.Errorf("error %v", err))
			os.Exit(1)
		}

		tweetText, _ := cmd.Flags().GetString("text")

		tweet, err := internal.CreateTweet(twitterClient, tweetText)
		if err != nil {
			fmt.Println(fmt.Errorf("error %v", err))
			os.Exit(1)
		}

		fmt.Printf("tweet: %v", tweet)
	},
}
