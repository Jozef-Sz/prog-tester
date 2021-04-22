package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	// testCases := testcase.GetTestCasesFromSchema(testSchemaName + ".xml")
	// for _, testcase := range testCases {
	// 	testcase.evaluate(discardOutput)
	// }

	// fmt.Printf("Running test with %s\n", args[0])
	// fmt.Printf("Save output? %t\n", generateOutput)
}

func genCommand(cmd *cobra.Command, args []string) {
	count := 0
	if len(args) != 0 {
		count, _ = strconv.Atoi(args[0])
	}
	fmt.Println("Sorry, not implemented yet.", count)
}
