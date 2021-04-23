package tester

import (
	"fmt"
	"tester/testcase"
)

const MAX_INPUT_LEN = 50

func printResult(result testcase.Result) {
	fmt.Printf("TestCase%d..........", result.TestCase.ID)
	if checkResultSuccess(&result) == 0 {
		fmt.Print("FAIL\n")
		printResultMetadata(result)
	} else {
		fmt.Print("OK\n")
	}
}

func printResultMetadata(result testcase.Result) {
	fmt.Printf("  args: %v\n", result.ArgsUsed)
	if len(result.Output) > MAX_INPUT_LEN {
		fmt.Printf("  input: %s....\n", result.Output[:50])
	} else {
		fmt.Printf("  input: %s\n", result.Output)
	}
}
