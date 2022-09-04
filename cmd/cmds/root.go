package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "go-twitter",
		Short: "go-twitter is a tool to interact with twitter api through command-line.",
		Long:  `You need to provide 4 keys, a consumer key and its secret, and an access token and its secret.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
