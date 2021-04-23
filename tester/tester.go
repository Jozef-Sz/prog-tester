package tester

import (
	"fmt"
	"strings"
	"tester/testcase"
)

type ResultSuccess int

const (
	SUCCESSFUL        ResultSuccess = 0
	OUTPUT_MISMATCH                 = 1
	EXITCODE_MISMATCH               = 2
)

func hasFlag(success ResultSuccess, flag ResultSuccess) bool {
	return (success & flag) != 0
}

type Tester struct {
	passed  int
	results []testcase.Result
}

func NewTester(testCases []testcase.TestCase, executable string) Tester {
	tester := Tester{passed: 0}
	for _, testcase := range testCases {
		result := testcase.Run(executable)
		tester.results = append(tester.results, result)

		if checkResultSuccess(&result) == SUCCESSFUL {
			tester.passed++
		}
	}
	return tester
}

func checkResultSuccess(result *testcase.Result) ResultSuccess {
	var success ResultSuccess = SUCCESSFUL
	// TODO: during comparison you have to handle \r\n == \n (windows stuff)
	if result.TestCase.Expect != strings.TrimRight(result.Output, " ") {
		fmt.Println("Whaaaattt it does not equals")
		success |= OUTPUT_MISMATCH
	}
	if result.TestCase.Exitcode != result.ExitCode {
		success |= EXITCODE_MISMATCH
	}
	return success
}

func (tester *Tester) EvaluateResults(discardOutput bool) {
	for _, result := range tester.results {
		printTestcaseResult(result)
		if !discardOutput {
			saveResultToFile(result)
		}
	}
	printTestConslusion(tester)
}
