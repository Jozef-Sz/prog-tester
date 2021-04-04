package cli

import (
	"fmt"
	"strconv"
	"tester/xml"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run [exe name]",
	Short: "run tester with the specified executable",
	Args:  cobra.ExactArgs(1),
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	testCases := xml.ParseXmlSchema(testSchemaName + ".xml")
	fmt.Println("Running test with", testCases)
	fmt.Printf("Running test with %s\n", args[0])
	fmt.Printf("Save output? %t\n", generateOutput)
	fmt.Printf("Input %s\n", testSchemaName)
}

var GenCmd = &cobra.Command{
	Use:   "gen [?count]",
	Short: "generate required files and folders",
	Args:  cobra.MaximumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			_, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("Supplied value %s is not a number\n", args[0])
			}
		}
		return nil
	},
	Run: genCommand,
}

func genCommand(cmd *cobra.Command, args []string) {
	count := 0
	if len(args) != 0 {
		count, _ = strconv.Atoi(args[0])
	}
	fmt.Println("Sorry, not implemented yet.", count)
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print Tester version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tester v%s\n", VERSION)
	},
}
