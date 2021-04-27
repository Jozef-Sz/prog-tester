package cli

import (
	"fmt"
	"strconv"
	"tester/testcase"
	"tester/tester"

	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	// TODO: handle file extensions .xml and .exe (args[0])
	// so if user enters file.xml do not hard code .xml to the end
	testCases := testcase.GetTestCasesFromSchema(testSchemaName + ".xml")
	tester := tester.RunTester(testCases, args[0])
	tester.EvaluateResults(discardOutput)
}

func genCommand(cmd *cobra.Command, args []string) {
	count := 0
	if len(args) != 0 {
		count, _ = strconv.Atoi(args[0])
	}
	fmt.Println("Sorry, not implemented yet.", count)
}
