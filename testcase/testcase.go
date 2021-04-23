package testcase

import (
	"bytes"
	"fmt"
	"os/exec"
	"tester/platform"
)

type TestCase struct {
	ID       int
	Args     string
	Input    string
	Expect   string
	Exitcode int
}

type Result struct {
	Output   string
	ExitCode int
	ErrorMsg string
	TestCase TestCase
	// redundant property, used only for debug
	// and for now, it's also used durin printing out results
	// to make sure the args were split correctly
	ArgsUsed []string
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
	inputBytesBuffer.Write([]byte(testcase.Input))
	fmt.Printf("%s\n", testcase.Input)
	fmt.Println("-----------")
	args := splitArguments(testcase.Args)
	cmd := exec.Command(platform.GetExeName(exe), args...)
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
