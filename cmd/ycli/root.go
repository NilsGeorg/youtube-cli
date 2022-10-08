package ycli

import (
	"encoding/json"
	"os"

	"github.com/NilsGeorg/ycli/pkg/config"
	"github.com/NilsGeorg/ycli/pkg/youtube"

	"github.com/spf13/cobra"
)

var YoutubeClient youtube.Client

type PrintMessage struct {
	Content interface{}
	Type    string
}

var RootCmd = &cobra.Command{
	Use:     "youtube-cli",
	Aliases: []string{"youtube"},
	Version: config.Version,
	Short:   "CLI for calling the youtube API",
	RunE: func(cmd *cobra.Command, args []string) error {

		return cmd.Help()
	},
}

func PrettyPrint(cmd *cobra.Command, i interface{}) error {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return err
	}

	cmd.Println(string(s))

	return nil
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		RootCmd.PrintErrf("ERROR: %s", err)

		os.Exit(1)
	}
}

func init() {
	YoutubeClient = youtube.DefaultClient
}
