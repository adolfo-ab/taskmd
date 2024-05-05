package cmd

import (
	"fmt"
	"github.com/adolfo-ab/taskmd/pkg/util"
	"github.com/adolfo-ab/taskmd/pkg/util/entities"
	"github.com/spf13/cobra"
	"os"
)

var pendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "return all pending tasks in a given directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := util.VerifyPathExists(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pending, err := util.GetPendingTasks(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create map of tasks by file
		for _, tf := range pending {
			fmt.Printf("- %s:\n", tf.Path)
			for _, task := range tf.Tasks {
				fmt.Printf("%s%s%s%s\n", entities.Pending, util.RedColor, task.Content, util.ResetColor)
			}
		}
		fmt.Printf("Total number of pending tasks in %s: %d\n", path, util.GetTotalNumberOfTasks(pending))

	},
}

func init() {
	rootCmd.AddCommand(pendingCmd)
}
