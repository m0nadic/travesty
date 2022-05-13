package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"travesty/internal/app/util"
)

const defaultFile = "travesty.yaml"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "displays version information",
	Long:  "Display version information and exit.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "creates a config file",
	Long:  "Creates a new Travesty.yaml in the current folder.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Generating", defaultFile)
			err := util.GenerateFile(defaultFile)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				return
			}
			return
		}

		for _, arg := range args {
			fmt.Println("Generating", arg)
			err := util.GenerateFile(arg)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				continue
			}
		}

	},
}
