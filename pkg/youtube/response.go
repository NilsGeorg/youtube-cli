package youtube

import (
	"time"
)

type ResponseType interface {
	SearchResource | VideoResource | VideoCategoryResource | ChannelResource
}

type Response[T ResponseType] struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PrevPageToken string `json:"prevPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		ResultsPerPage int64 `json:"resultsPerPage"`
		TotalResults   int64 `json:"totalResults"`
	} `json:"pageInfo"`
	Items []*T `json:"items"`
}

type SearchResource struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	ID   struct {
		Kind    string `json:"kind"`
		VideoID string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		PublishedAt time.Time `json:"publishedAt"`
		ChannelID   string    `json:"channelId"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		ChannelTitle         string    `json:"channelTitle"`
		LiveBroadcastContent string    `json:"liveBroadcastContent"`
		PublishTime          time.Time `json:"publishTime"`
	} `json:"snippet"`
}

type VideoCategoryResource struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		ChannelID  string `json:"channelId"`
		Title      string `json:"title"`
		Assignable bool   `json:"assignable"`
	} `json:"snippet"`
}

type VideoResource struct {
	Etag string `json:"etag"`
	ID   string `json:"id"`
	Kind string `json:"kind"`

	ContentDetails struct {
		Caption         string   `json:"caption"`
		ContentRating   struct{} `json:"contentRating"`
		Definition      string   `json:"definition"`
		Dimension       string   `json:"dimension"`
		Duration        string   `json:"duration"`
		LicensedContent bool     `json:"licensedContent"`
		Projection      string   `json:"projection"`
	} `json:"contentDetails"`

	Statistics struct {
		ViewCount     string `json:"viewCount"`
		LikeCount     string `json:"likeCount"`
		FavoriteCount string `json:"favoriteCount"`
		CommentCount  string `json:"commentCount"`
	} `json:"statistics"`

	Snippet struct {
		CategoryID           string `json:"categoryId"`
		ChannelID            string `json:"channelId"`
		ChannelTitle         string `json:"channelTitle"`
		DefaultAudioLanguage string `json:"defaultAudioLanguage"`
		DefaultLanguage      string `json:"defaultLanguage"`
		Description          string `json:"description"`
		LiveBroadcastContent string `json:"liveBroadcastContent"`
		Localized            struct {
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"localized"`
		PublishedAt string   `json:"publishedAt"`
		Tags        []string `json:"tags"`
		Thumbnails  struct {
			Default struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"default"`
			High struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"high"`
			Maxres struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"maxres"`
			Medium struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"medium"`
			Standard struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"standard"`
		} `json:"thumbnails"`
		Title string `json:"title"`
	} `json:"snippet"`
}

type ChannelResource struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CustomURL   string    `json:"customUrl"`
		PublishedAt time.Time `json:"publishedAt"`
		Country     string    `json:"country"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		Localized struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"localized"`
	} `json:"snippet"`
	ContentDetails struct {
		RelatedPlaylists struct {
			Likes   string `json:"likes"`
			Uploads string `json:"uploads"`
		} `json:"relatedPlaylists"`
	} `json:"contentDetails"`
	Statistics struct {
		ViewCount             string `json:"viewCount"`
		SubscriberCount       string `json:"subscriberCount"`
		HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
		VideoCount            string `json:"videoCount"`
	} `json:"statistics"`
	TopicDetails struct {
		TopicIds        []string `json:"topicIds"`
		TopicCategories []string `json:"topicCategories"`
	} `json:"topicDetails"`
	Status struct {
		PrivacyStatus     string `json:"privacyStatus"`
		IsLinked          bool   `json:"isLinked"`
		LongUploadsStatus string `json:"longUploadsStatus"`
		MadeForKids       bool   `json:"madeForKids"`
	} `json:"status"`
	BrandingSettings struct {
		Channel struct {
			Title                      string `json:"title"`
			Description                string `json:"description"`
			Keywords                   string `json:"keywords"`
			TrackingAnalyticsAccountID string `json:"trackingAnalyticsAccountId"`
			UnsubscribedTrailer        string `json:"unsubscribedTrailer"`
			Country                    string `json:"country"`
		} `json:"channel"`
		Image struct {
			BannerExternalURL string `json:"bannerExternalUrl"`
		} `json:"image"`
	} `json:"brandingSettings"`
}
