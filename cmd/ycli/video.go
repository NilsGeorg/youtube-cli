package ycli

import (
	"github.com/spf13/cobra"
)

var videoCmd = &cobra.Command{
	Use: "video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var getVideo = &cobra.Command{
	Use:  "get",
	Args: cobra.RangeArgs(1, 50),
	RunE: func(cmd *cobra.Command, args []string) error {
		videos, err := YoutubeClient.Videos(args)
		if err != nil {
			return err
		}

		return PrettyPrint(cmd, videos.Items)
	},
}

var searchVideo = &cobra.Command{
	Use: "search",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword, _ := cmd.Flags().GetString("keyword")
		category, _ := cmd.Flags().GetString("category")
		region, _ := cmd.Flags().GetString("region")

		videos, err := YoutubeClient.VideoSearch(keyword, category, region)
		if err != nil {
			return err
		}

		return PrettyPrint(cmd, videos.Items)
	},
}

var getVideoChart = &cobra.Command{
	Use: "most-popular",
	RunE: func(cmd *cobra.Command, args []string) error {
		category, _ := cmd.Flags().GetString("category")
		region, _ := cmd.Flags().GetString("region")

		videos, err := YoutubeClient.VideosByChart(region, category)

		if err != nil {
			return err
		}

		return PrettyPrint(cmd, videos.Items)
	},
}

var getVideoCategories = &cobra.Command{
	Use: "categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		region, _ := cmd.Flags().GetString("region")
		category, _ := cmd.Flags().GetString("category")

		categories, err := YoutubeClient.VideoGetCategories(region, category)
		if err != nil {
			return err
		}

		return PrettyPrint(cmd, categories.Items)
	},
}

func init() {
	RootCmd.AddCommand(videoCmd)
	videoCmd.AddCommand(getVideo, searchVideo, getVideoChart, getVideoCategories)

	searchVideo.Flags().StringP("keyword", "k", "", "The keyword to search for")
	_ = searchVideo.MarkFlagRequired("keyword")
	searchVideo.Flags().StringP("category", "c", "", "Category ID to search in (Optional)")
	searchVideo.Flags().StringP("region", "r", "", "Only search in specified region (Optional)")

	getVideoCategories.Flags().StringP("region", "r", "", "Only search in specified region")
	getVideoCategories.Flags().StringP("category", "c", "", "Category ID to search in (Optional)")
	_ = getVideoCategories.MarkFlagRequired("region")

	getVideoChart.Flags().StringP("region", "r", "", "Only search in specified region (Optional)")
	getVideoChart.Flags().StringP("category", "c", "", "Category ID to search in (Optional)")
}
