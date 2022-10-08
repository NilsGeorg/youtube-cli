package ycli

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/NilsGeorg/ycli/pkg/youtube"
	"github.com/stretchr/testify/assert"
)

func Test_Root(t *testing.T) {
	result := test.CmdTestRunner(t, RootCmd.Root(), test.NewCmdArgs())

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func Test_PrettyPrint(t *testing.T) {
	var videoResponse youtube.Response[youtube.VideoResource]
	marshalErr := json.Unmarshal([]byte(mock.VideoResponse), &videoResponse)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}

	successOutput := bytes.NewBufferString("")
	errorOutput := bytes.NewBufferString("")

	RootCmd.SetOut(successOutput)
	RootCmd.SetErr(errorOutput)

	prettyPrintErr := PrettyPrint(RootCmd, videoResponse)
	assert.NoError(t, prettyPrintErr)

	stdout, err := io.ReadAll(successOutput)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	stderr, err := io.ReadAll(errorOutput)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	assert.Empty(t, stderr)
	assert.JSONEq(t, mock.VideoResponse, string(stdout))
}
