package ycli

import (
	"net/http"
	"os"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/pkg/youtube"
)

func TestMain(m *testing.M) {
	YoutubeClient = youtube.Client{
		HTTPClient:     &http.Client{},
		Authentication: mock.Authentication{},
		YoutubeAPIURL:  "youtube-cli.local",
		YoutubeAPIPath: "/api",
	}

	code := m.Run()
	os.Exit(code)
}
