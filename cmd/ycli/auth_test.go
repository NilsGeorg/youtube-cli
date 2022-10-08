package ycli

import (
	"fmt"
	"testing"

	"github.com/NilsGeorg/ycli/internal/mock"
	"github.com/NilsGeorg/ycli/internal/test"
	"github.com/stretchr/testify/assert"
)

func Test_AuthCmd(t *testing.T) {
	args := test.NewCmdArgs().AddArg("auth")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func Test_TokenCmd(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg("auth").
		AddArg("token")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func Test_TokenSet(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg("auth").
		AddArg("token").
		AddArg("set").
		AddArg(mock.AuthenticationToken)

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
	assert.Equal(t, "Token was set successfully\n", result.Stdout)
}

func Test_TokenSet_Non(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg("auth").
		AddArg("token").
		AddArg("set")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts 1 arg(s), received 0\n", result.Stderr)
}

func Test_TokenSet_Multiple(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg("auth").
		AddArg("token").
		AddArg("set").
		AddArg(mock.AuthenticationToken).
		AddArg(mock.AuthenticationToken)

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.Error(t, result.Error)
	assert.NotEmpty(t, result.Stderr)
	assert.Equal(t, "Error: accepts 1 arg(s), received 2\n", result.Stderr)
}

func Test_TokenGet(t *testing.T) {
	args := test.NewCmdArgs().
		AddArg("auth").
		AddArg("token").
		AddArg("get")

	result := test.CmdTestRunner(t, RootCmd.Root(), args)

	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)
	assert.Equal(t, fmt.Sprintf("Token: %s\n", mock.AuthenticationToken), result.Stdout)
}
