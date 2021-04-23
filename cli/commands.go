package cli

import (
	"fmt"
	"strconv"
	"tester/testcase"
	"tester/tester"

	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	testCases := testcase.GetTestCasesFromSchema(testSchemaName + ".xml")
	tester := tester.NewTester(testCases, args[0])
	tester.EvaluateResults()
}

func genCommand(cmd *cobra.Command, args []string) {
	count := 0
	if len(args) != 0 {
		count, _ = strconv.Atoi(args[0])
	}
	fmt.Println("Sorry, not implemented yet.", count)
}
