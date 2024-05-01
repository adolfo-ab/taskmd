package taskmd

import (
	"fmt"
	"github.com/adolfo-ab/taskmd/pkg/taskmd"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "taskmd [arguments] [path]",
	Short: "taskmd is a simple CLI tool to provide task completion info.",
	Long:  `taskmd parses .md files in a given directory and provides metrics about task completion`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a path to a directory containing .md files")
			os.Exit(1)
		}

		path := args[0]

		err := taskmd.VerifyPathExists(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		percentage, err := taskmd.GetCompletionPercentage(args[0])
		if err != nil {
			return
		}

		colorCode := redColor // Default to red
		if percentage > 50 {
			colorCode = greenColor // Change to green if task completion is over 50%
		}

		fmt.Printf("Task completion ratio for tasks in %s: %s%.2f%%%s\n", path, colorCode, percentage, resetColor)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing taskmd CLI: %s'", err)
		os.Exit(1)
	}
}
