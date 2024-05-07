package cmd

import (
	"fmt"
	"github.com/adolfo-ab/taskmd/pkg/util"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "util [arguments] [path]",
	Short:   "util is a simple CLI tool to provide task completion info.",
	Long:    `util parses .md files in a given directory and provides metrics about task completion`,
	Args:    cobra.MinimumNArgs(1),
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := util.VerifyPathExists(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		report, err := util.GetCompletionReport(path)
		fmt.Println(report)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing util CLI: %s'", err)
		os.Exit(1)
	}
}
