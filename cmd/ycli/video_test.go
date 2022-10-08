package ycli

import (
	"fmt"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var testVideoOptions = struct {
	Command    string
	VideoPath  string
	SearchPath string
}{
	Command:    "video",
	VideoPath:  "/videos",
	SearchPath: "/search",
}

func Test_Video(t *testing.T) {
	args := test.NewCmdArgs().AddArg(testVideoOptions.Command)

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func Test_GetVideo(t *testing.T) {
	for _, videoIds := range testGetVideoIds() {
		var name = fmt.Sprintf("Amount-%d", len(videoIds))

		t.Run(name, func(t *testing.T) {
			result := runVideoByID(t, videoIds)

			assert.True(t, gock.IsDone())
			assert.NoError(t, result.Error)
			assert.Empty(t, result.Stderr)
			assert.JSONEq(t, mock.VideoResponseItems, result.Stdout)
		})
	}
}

func runVideoByID(t *testing.T, ids []string) test.CmdResult {
	t.Helper()

	defer gock.Off()
	gock.
		New("https://" + YoutubeClient.YoutubeAPIURL).
		Get(YoutubeClient.YoutubeAPIPath + testVideoOptions.VideoPath).
		Reply(200).
		JSON(mock.VideoResponse)

	args := test.NewCmdArgs().
		AddArg(testVideoOptions.Command).
		AddArg("get")

	for _, channelID := range ids {
		args = args.AddArg(channelID)
	}

	return test.CmdTestRunner(t, RootCmd.Root(), args)
}

func testGetVideoIds() [][]string {
	var rickID = "dQw4w9WgXcQ"
	var videoTest = [][]string{}
	var amountTestRuns = 50

	for i := 1; i <= amountTestRuns; i++ {
		channelIds := []string{rickID}
		for j := 2; j <= i; j++ {
			channelIds = append(channelIds, fmt.Sprintf("%s-%d", rickID, i))
		}

		videoTest = append(videoTest, channelIds)
	}

	return videoTest
}

func Test_GetVideo_MoreThan50(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testVideoOptions.Command).
		AddArg("get")

	for i := 1; i <= 51; i++ {
		args = args.AddArg(fmt.Sprintf("videoId-%d", i))
	}

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts between 1 and 50 arg(s), received 51\n", result.Stderr)
}

func Test_GetVideo_NoId(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testVideoOptions.Command).
		AddArg("get")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts between 1 and 50 arg(s), received 0\n", result.Stderr)
}

var testSearchVideoMatrix = []struct {
	Keyword   string
	Category  string
	Region    string
	IsSuccess bool
}{
	{"rick astley", "US", "10", true},
	{"rick astley", "US", "", true},
	{"rick astley", "", "10", true},
	{"rick astley", "", "", true},
	{"", "US", "10", false},
	{"", "US", "", false},
	{"", "", "10", false},
	{"", "", "", false},
}

func Test_SearchVideo(t *testing.T) {
	for _, testCase := range testSearchVideoMatrix {
		var name = fmt.Sprintf("%s-%s-%s", testCase.Keyword, testCase.Category, testCase.Region)

		t.Run(name, func(t *testing.T) {
			if testCase.IsSuccess {
				defer gock.Off()
				gock.
					New("https://" + YoutubeClient.YoutubeAPIURL).
					Get(YoutubeClient.YoutubeAPIPath + testVideoOptions.SearchPath).
					Reply(200).
					JSON(mock.VideoSearchResponse)
			}

			args := test.NewCmdArgs().
				AddArg(testVideoOptions.Command).
				AddArg("search")

			if testCase.Keyword != "" {
				args = args.AddFlag("keyword", testCase.Keyword, false)
			}

			if testCase.Category != "" {
				args = args.AddFlag("category", testCase.Category, false)
			}

			if testCase.Region != "" {
				args = args.AddFlag("region", testCase.Region, false)
			}

			result := test.CmdTestRunner(t, RootCmd.Root(), args)

			if testCase.IsSuccess {
				assert.True(t, gock.IsDone())
				assert.NoError(t, result.Error)
				assert.Empty(t, result.Stderr)
				assert.JSONEq(t, mock.VideoSearchResponseItems, result.Stdout)
			} else {
				assert.Error(t, result.Error)
				assert.NotEmpty(t, result.Stderr)
				assert.NotEmpty(t, result.Stdout)
			}
		})
	}

}
