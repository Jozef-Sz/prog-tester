package tester

import (
	"strings"
	"tester/testcase"
)

type Tester struct {
	passed  int
	results []testcase.Result
}

func NewTester(testCases []testcase.TestCase, executable string) Tester {
	tester := Tester{passed: 0}
	for i, testcase := range testCases {
		result := testcase.Run(executable)
		tester.results[i] = result
		tester.passed += checkResultSuccess(&result)
	}
	return tester
}

func checkResultSuccess(result *testcase.Result) int {
	if result.Output == "" || result.ErrorMsg != "" {
		return 0
	}
	if result.TestCase.Expect != strings.TrimRight(result.Output, " ") {
		return 0
	}
	if result.TestCase.Exitcode != result.ExitCode {
		return 0
	}
	return 1
}

func (tester *Tester) EvaluateResults(discardOutput bool) {
	for _, result := range tester.results {
		printResult(result)
		if !discardOutput {
			// save output and exepected to txt file
		}
	}
}
