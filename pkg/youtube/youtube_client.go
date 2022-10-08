package youtube

import (
	"net/http"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/pkg/config"
)

type Client struct {
	HTTPClient     HTTPClient
	Authentication config.Authentication
	YoutubeAPIURL  string
	YoutubeAPIPath string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var DefaultClient = Client{
	Authentication: config.AuthenticationImpl{},
	HTTPClient:     http.DefaultClient,
	YoutubeAPIURL:  config.YoutubeAPIURL,
	YoutubeAPIPath: config.YoutubeAPIPath,
}

var TestHTTPClient = Client{
	Authentication: mock.Authentication{},
	HTTPClient:     http.DefaultClient,
	YoutubeAPIURL:  "youtube-cli.local",
	YoutubeAPIPath: "/api",
}

func TestClient(jsonResponse string) Client {
	return Client{
		HTTPClient:     mock.HTTPClient{JSON: jsonResponse},
		Authentication: mock.Authentication{},
		YoutubeAPIURL:  "youtube-cli.local",
		YoutubeAPIPath: "/api",
	}
}
