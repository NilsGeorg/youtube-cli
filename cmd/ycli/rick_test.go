package ycli

import (
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var rickCommand = "rick"

func Test_RickCmd(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(rickCommand).
		AddArg("classic")

	defer gock.Off()
	gock.
		New("https://" + YoutubeClient.YoutubeAPIURL).
		Get(YoutubeClient.YoutubeAPIPath + "/search").
		Reply(200).
		JSON(mock.VideoSearchResponse)

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)
	assert.Equal(t, "https://www.youtube.com/watch?v=dQw4w9WgXcQ\n", result.Stdout)
}

func Test_RickClassic(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg(rickCommand).
		AddArg("classic")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)
	assert.Equal(t, "https://www.youtube.com/watch?v=dQw4w9WgXcQ\n", result.Stdout)
}
