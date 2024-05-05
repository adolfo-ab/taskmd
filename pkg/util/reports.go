package util

import (
	"fmt"
	"github.com/adolfo-ab/taskmd/pkg/util/entities"
)

func GetCompletionReport(path string) (string, error) {
	tfs, err := getTaskFiles(path)
	if err != nil {
		return "", err
	}

	percentage := GetCompletionPercentage(tfs)
	average := GetAverageNumberOfTasksPerFile(tfs)

	return fmt.Sprintf("--------Task Completion Report--------\n"+
		"Directory: %s\n"+
		"Task Completion Percentage: %.2f\n"+
		"Average Number of Tasks per File: %.2f\n",
		path, percentage, average), nil
}

func GetPendingTasks(path string) ([]entities.TaskFile, error) {
	tf, err := getTaskFiles(path)
	if err != nil {
		return nil, err
	}

	return filterPendingTasks(tf), nil
}
