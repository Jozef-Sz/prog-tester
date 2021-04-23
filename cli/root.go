package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tester",
	Short: "Tester is a tool for testing executable files by providing input and comparing it's output",
}

const DEFAULT_SCHEMA_NAME = "test_schema"

var (
	discardOutput  bool
	testSchemaName string
)

func Execute(version string) {
	initCLI(version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initCLI(version string) {
	addFlags()
	rootCmd.AddCommand(RunCmd)
	rootCmd.AddCommand(GenCmd)
	rootCmd.AddCommand(VersionCmd(version))
}

func addFlags() {
	RunCmd.Flags().BoolVarP(&discardOutput, "discard-output", "d", false, "Do not save the output to txt files")
	RunCmd.Flags().StringVarP(&testSchemaName, "name", "n", DEFAULT_SCHEMA_NAME, "Xml file name of test cases")
}
