package test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/NilsGeorg/ycli/pkg/config"
	"github.com/spf13/cobra"
)

type CmdArgs struct {
	Args   []string
	Config config.Authentication
}

type CmdResult struct {
	Stdout string
	Stderr string
	Error  error
}

func NewCmdArgs() CmdArgs {
	return CmdArgs{
		Args: []string{},
	}
}

func (cmdArgs CmdArgs) AddArg(arg string) CmdArgs {
	cmdArgs.Args = append(cmdArgs.Args, arg)
	return cmdArgs
}

func (cmdArgs CmdArgs) AddFlag(flag, value string, shorthand bool) CmdArgs {
	value = strings.TrimSpace(value)

	var flagIndicator string
	if shorthand {
		flagIndicator = "-"
	} else {
		flagIndicator = "--"
	}

	cmdArgs.Args = append(cmdArgs.Args, fmt.Sprintf(`%s%s`, flagIndicator, flag))
	if value != "" {
		cmdArgs.Args = append(cmdArgs.Args, value)
	}

	return cmdArgs
}

func CmdTestRunner(t *testing.T, cmd *cobra.Command, cmdArgs CmdArgs) CmdResult {
	t.Helper()

	successOutput := bytes.NewBufferString("")
	errorOutput := bytes.NewBufferString("")

	cmd.SetOut(successOutput)
	cmd.SetErr(errorOutput)

	cmd.SetArgs(cmdArgs.Args)
	cmdError := cmd.Execute()

	stdout, err := io.ReadAll(successOutput)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	stderr, err := io.ReadAll(errorOutput)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	return CmdResult{
		Stdout: string(stdout),
		Stderr: string(stderr),
		Error:  cmdError,
	}
}
