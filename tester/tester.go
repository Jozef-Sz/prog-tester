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
	if result.output == "" || result.errorMsg != "" {
		return 0
	}
	if result.testCase.expect != strings.TrimRight(result.output, " ") {
		return 0
	}
	if result.testCase.exitcode != result.exitCode {
		return 0
	}
	return 1
}

func (tester *Tester) EvaluateResults() {}
