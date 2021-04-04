package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tester",
	Short: "Tester is a tool for testing executable files by providing input and comparing it's output",
}

const DEFAULT_SCHEMA_NAME = "test_schema"

var (
	generateOutput bool
	testSchemaName string
)

func Execute() {
	initCLI()
	rootCmd.Execute()
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
}

func initCLI() {
	addFlags()
	rootCmd.AddCommand(RunCmd)
	rootCmd.AddCommand(GenCmd)
	rootCmd.AddCommand(VersionCmd)
}

func addFlags() {
	RunCmd.Flags().BoolVarP(&generateOutput, "save-output", "s", false, "Save test case's output to txt")
	RunCmd.Flags().StringVarP(&testSchemaName, "name", "n", DEFAULT_SCHEMA_NAME, "Xml file name of test cases")
}
