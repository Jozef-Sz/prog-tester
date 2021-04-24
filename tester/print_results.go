package tester

import (
	"fmt"
	"strings"
	"tester/testcase"
)

const MAX_INPUT_LEN = 50

func printTestcaseResult(result testcase.Result) {
	fmt.Printf("TestCase%d..........", result.TestCase.ID)
	// fmt.Printf("%v", result)
	resultSuccess := checkResultSuccess(&result)
	if resultSuccess == 0 {
		fmt.Print("OK\n")
	} else {
		fmt.Print("FAIL\n")
		printTestcaseFailures(result, resultSuccess)
	}
}

func printTestcaseFailures(result testcase.Result, success ResultSuccess) {
	printTestcaseInputs(result)

	if hasFlag(success, OUTPUT_MISMATCH) {
		fmt.Println("  OUTPUT: not matching with expected")
	}
	if hasFlag(success, EXITCODE_MISMATCH) {
		fmt.Printf("  EXIT_CODE: got: %d, expected: %d\n", result.ExitCode, result.TestCase.Exitcode)
	}
	if result.ErrorMsg != "" {
		fmt.Printf("  PROGRAM_EXITED_WITH: %s\n", result.ErrorMsg)
	}
}

func printTestcaseInputs(result testcase.Result) {
	// TODO: check also for new lines, if there is cut the string and put there ....
	fmt.Printf("  args: %v\n", result.ArgsUsed)
	if len(result.Output) > MAX_INPUT_LEN {
		fmt.Printf("  input: %s....\n", result.TestCase.Input[:50])
	} else {
		fmt.Printf("  input: %s\n", result.TestCase.Input)
	}
}

func printTestConslusion(tester *Tester) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Ran %d tests ", len(tester.results))
	if tester.passed == len(tester.results) {
		fmt.Println("\nTEST PASSED")
	} else {
		fmt.Printf("(passed: %d, failed: %d)", tester.passed, len(tester.results)-tester.passed)
		fmt.Println("\nTEST FAILED")
	}
}
