package ycli

import (
	"github.com/spf13/cobra"
)

var channelCmd = &cobra.Command{
	Use: "channel",
	RunE: func(cmd *cobra.Command, args []string) error {

		return cmd.Help()
	},
}

var channelByID = &cobra.Command{
	Use:  "id",
	Args: cobra.RangeArgs(1, 50),
	RunE: func(cmd *cobra.Command, args []string) error {
		channels, err := YoutubeClient.Channels(args)
		if err != nil {
			return err
		}

		return PrettyPrint(cmd, channels.Items)
	},
}

var channelByUsername = &cobra.Command{
	Use:  "username",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		channels, err := YoutubeClient.ChannelByUsername(args[0])
		if err != nil {
			return err
		}

		return PrettyPrint(cmd, channels.Items)
	},
}

func init() {
	RootCmd.AddCommand(channelCmd)
	channelCmd.AddCommand(channelByID, channelByUsername)
}
