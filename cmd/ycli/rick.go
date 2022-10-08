package ycli

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var rickCmd = &cobra.Command{
	Use:   "rick",
	Short: "Returns random rick roll",
	RunE: func(cmd *cobra.Command, args []string) error {
		ricks, err := YoutubeClient.VideoSearch("Never gonna give you up", "", "")
		if err != nil {
			return err
		}

		rand.Seed(time.Now().UnixNano())

		randomIndex := rand.Intn(len(ricks.Items) - 1) //nolint:gosec // not used for crypto
		randomRick := ricks.Items[randomIndex]

		cmd.Printf("https://www.youtube.com/watch?v=%s", randomRick.ID.VideoID)
		return nil
	},
}

var classic = &cobra.Command{
	Use:   "classic",
	Short: "returns the original never gonna give you up",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("https://www.youtube.com/watch?v=dQw4w9WgXcQ\n")
	},
}

func init() {
	RootCmd.AddCommand(rickCmd)
	rickCmd.AddCommand(classic)
}
