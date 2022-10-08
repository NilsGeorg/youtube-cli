package youtube

import "strings"

var channelPart = "brandingSettings,contentDetails,id,localizations,snippet,statistics,status,topicDetails"

func (c Client) Channels(ids []string) (*Response[ChannelResource], error) {
	joinedIds := strings.Join(ids, ",")

	var request = NewRequest().
		AddParam("id", joinedIds).
		AddParam("part", channelPart).
		AddParam("maxResults", "50")

	return NewAPI[ChannelResource](c).Get("/channels", request)
}

func (c Client) ChannelByUsername(username string) (*Response[ChannelResource], error) {
	var request = NewRequest().
		AddParam("forUsername", username).
		AddParam("part", channelPart)

	return NewAPI[ChannelResource](c).Get("/channels", request)
}
