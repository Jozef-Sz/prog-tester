package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	// testCases := xml.ParseXmlSchema(testSchemaName + ".xml")
	fmt.Printf("Running test with %s\n", args[0])
	fmt.Printf("Save output? %t\n", generateOutput)
	// fmt.Printf("Input %s\n", testSchemaName)
}

func genCommand(cmd *cobra.Command, args []string) {
	count := 0
	if len(args) != 0 {
		count, _ = strconv.Atoi(args[0])
	}
	fmt.Println("Sorry, not implemented yet.", count)
}
