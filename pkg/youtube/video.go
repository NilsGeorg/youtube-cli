package youtube

import (
	"errors"
	"strings"
)

func (c Client) VideoSearch(keyword, category, region string) (*Response[SearchResource], error) {
	var request = NewRequest().
		AddParam("part", "snippet").
		AddParam("q", keyword).
		AddParam("category", category).
		AddParam("region", region).
		AddParam("maxResults", "50")

	return NewAPI[SearchResource](c).Get("/search", request)
}

func (c Client) Videos(ids []string) (*Response[VideoResource], error) {
	joinedIds := strings.Join(ids, ",")

	var request = NewRequest().
		AddParam("id", joinedIds).
		AddParam("part", "snippet,contentDetails,statistics")

	return NewAPI[VideoResource](c).Get("/videos", request)
}

func (c Client) VideosByChart(region, category string) (*Response[VideoResource], error) {
	var request = NewRequest().
		AddParam("region", region).
		AddParam("category", category).
		AddParam("chart", "mostPopular"). // currently the only accessible chart
		AddParam("part", "snippet,contentDetails,statistics")

	return NewAPI[VideoResource](c).Get("/videos", request)
}

func (c Client) VideoGetCategories(region, category string) (*Response[VideoCategoryResource], error) {
	if region != "" && category != "" {
		return nil, errors.New("you can only provide one of the params: region, category")
	}

	var request = NewRequest().
		AddParam("part", "snippet").
		AddParam("id", category).
		AddParam("regionCode", region)

	return NewAPI[VideoCategoryResource](c).Get("/videoCategories", request)
}
