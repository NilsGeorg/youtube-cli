package ycli

import (
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use: "auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
var tokenCmd = &cobra.Command{
	Use: "token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var setAuthToken = &cobra.Command{
	Use:  "set",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := YoutubeClient.Authentication.SetToken(args[0])
		if err != nil {
			return err
		}

		cmd.Println("Token was set successfully")

		return nil
	},
}

var getAuthToken = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := YoutubeClient.Authentication.GetToken()
		if err != nil {
			return err
		}

		cmd.Printf("Token: %s\n", token)

		return nil
	},
}

var deleteAuthToken = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := YoutubeClient.Authentication.SetToken("")
		if err != nil {
			return err
		}

		cmd.Println("Token was deleted successfully")

		return nil
	},
}

func init() {
	RootCmd.AddCommand(authCmd)
	authCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(setAuthToken, getAuthToken, deleteAuthToken)
}
