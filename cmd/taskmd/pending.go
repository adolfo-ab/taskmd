package taskmd

import (
	"fmt"
	"github.com/adolfo-ab/taskmd/pkg/taskmd"
	"github.com/spf13/cobra"
	"os"
)

var pendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "return all pending tasks in a given directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := taskmd.VerifyPathExists(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pendingTasks, err := taskmd.GetPendingTasks(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create map of tasks by file
		taskByFile := make(map[string][]taskmd.Task)
		for _, task := range pendingTasks {
			taskByFile[task.File] = append(taskByFile[task.File], task)
		}

		for file, tasks := range taskByFile {
			fmt.Printf("- %s:\n", file)
			for _, task := range tasks {
				fmt.Printf("  - [ ] %s%s%s\n", redColor, task.Content, resetColor)
			}
		}
		fmt.Printf("Total number of pending tasks in %s: %d\n", path, len(pendingTasks))

	},
}

func init() {
	rootCmd.AddCommand(pendingCmd)
}
