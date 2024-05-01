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
		fmt.Println("Pending tasks:")
		for _, task := range pendingTasks {
			fmt.Printf("- [ ] %s\n", task.Content)
		}
		fmt.Printf("Total number of pending tasks in %s: %d\n", path, len(pendingTasks))

	},
}

func init() {
	rootCmd.AddCommand(pendingCmd)
}
