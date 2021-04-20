package cli

import (
	"fmt"
	"strconv"
	"tester/comparator"
	"tester/xml"

	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	testCases := xml.ParseXmlSchema(testSchemaName + ".xml")
	// for _, testCase := range testCases {
	// 	fmt.Printf("\"%s\"\n", testCase.Input)
	// }

	comparator.RunTesting(testCases, args[0], generateOutput)
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
