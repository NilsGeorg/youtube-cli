package ycli

import (
	"fmt"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var testChannelOptions = struct {
	Command string
	URLPath string
}{
	Command: "channel",
	URLPath: "/channels",
}

func Test_Channel(t *testing.T) {
	args := test.NewCmdArgs().AddArg(testChannelOptions.Command)

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func Test_ChannelById(t *testing.T) {
	for _, channelIDs := range testGetChannelIds() {
		var name = fmt.Sprintf("Amount-%d", len(channelIDs))

		t.Run(name, func(t *testing.T) {
			result := runChannelByID(t, channelIDs)

			assert.True(t, gock.IsDone())
			assert.NoError(t, result.Error)
			assert.Empty(t, result.Stderr)
			assert.JSONEq(t, mock.ChannelResponseItems, result.Stdout)
		})
	}
}

func runChannelByID(t *testing.T, ids []string) test.CmdResult {
	t.Helper()

	defer gock.Off()
	gock.
		New("https://" + YoutubeClient.YoutubeAPIURL).
		Get(YoutubeClient.YoutubeAPIPath + testChannelOptions.URLPath).
		Reply(200).
		JSON(mock.ChannelResponse)

	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("id")

	for _, channelID := range ids {
		args = args.AddArg(channelID)
	}

	return test.CmdTestRunner(t, RootCmd.Root(), args)
}

func testGetChannelIds() [][]string {
	var rickAstleyOfficialID = "UCuAXFkgsw1L7xaCfnd5JJOw"
	var channelTest = [][]string{}
	var amountTestRuns = 50

	for i := 1; i <= amountTestRuns; i++ {
		channelIds := []string{rickAstleyOfficialID}
		for j := 2; j <= i; j++ {
			channelIds = append(channelIds, fmt.Sprintf("%s-%d", rickAstleyOfficialID, i))
		}

		channelTest = append(channelTest, channelIds)
	}

	return channelTest
}

func Test_ChannelById_MoreThan50(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("id")

	for i := 1; i <= 51; i++ {
		args = args.AddArg(fmt.Sprintf("channelId-%d", i))
	}

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts between 1 and 50 arg(s), received 51\n", result.Stderr)
}

func Test_ChannelById_NoId(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("id")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts between 1 and 50 arg(s), received 0\n", result.Stderr)
}

func Test_ChannelByUsername(t *testing.T) {
	defer gock.Off()
	gock.
		New("https://" + YoutubeClient.YoutubeAPIURL).
		Get(YoutubeClient.YoutubeAPIPath + testChannelOptions.URLPath).
		Reply(200).
		JSON(mock.ChannelResponse)

	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("username").
		AddArg("RickAstleyOfficial")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.True(t, gock.IsDone())
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)
	assert.JSONEq(t, mock.ChannelResponseItems, result.Stdout)
}

func Test_ChannelByUsername_NoUsername(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("username")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", result.Stderr)
}

func Test_ChannelByUsername_MoreThanTwo(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(testChannelOptions.Command).
		AddArg("username").
		AddArg("RickAstleyOfficial").
		AddArg("RickAstleyInofficial")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts 1 arg(s), received 2\n", result.Stderr)
}
