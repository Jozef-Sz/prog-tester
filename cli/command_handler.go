package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run [exe name]",
	Short: "run tester with the specified executable",
	Args:  cobra.ExactArgs(1),
	Run:   runCommand,
}

var GenCmd = &cobra.Command{
	Use:   "gen [?count]",
	Short: "generate required files and folders",
	Args:  cobra.MaximumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			_, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("supplied value %s is not a number", args[0])
			}
		}
		return nil
	},
	Run: genCommand,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print Tester version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tester v%s\n", VERSION)
	},
}
