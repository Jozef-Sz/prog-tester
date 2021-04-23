package testcase

import (
	"bytes"
	"os/exec"
)

type TestCase struct {
	id       int
	args     string
	input    string
	expect   string
	exitcode int
}

type Result struct {
	output   string
	exitCode int
	errorMsg string
	testCase TestCase
	// redundant property, used only for debug
	// and for now, it's also used durin printing out results
	// to make sure the args were split correctly
	argsUsed []string
}

func (testcase *TestCase) Run(exe string) Result {
	command, argsUsed := createCommand(exe, testcase)
	output, err := command.Output()
	exitCode := command.ProcessState.ExitCode()
	if err != nil {
		return Result{string(output), exitCode, err.Error(), *testcase, argsUsed}
	}
	return Result{string(output), exitCode, "", *testcase, argsUsed}
}

func createCommand(exe string, testcase *TestCase) (*exec.Cmd, []string) {
	var inputBytesBuffer bytes.Buffer
	inputBytesBuffer.Write([]byte(testcase.input))
	args := splitArguments(testcase.args)
	cmd := exec.Command(exe, args...)
	cmd.Stdin = &inputBytesBuffer
	return cmd, args
}

func splitArguments(arguments string) []string {
	var arglist []string
	arg := ""
	insideQuotes := false
	for _, char := range arguments {
		if char == '"' {
			insideQuotes = !insideQuotes
			continue
		}
		if char == ' ' && !insideQuotes {
			arglist = append(arglist, arg)
			arg = ""
		} else {
			arg += string(char)
		}
	}
	return append(arglist, arg)
}
