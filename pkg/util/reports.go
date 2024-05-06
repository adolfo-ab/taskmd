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
	mean := GetAverageNumberOfTasksPerFile(tfs)
	std := GetStd(tfs, mean)
	maxTasks := GetMaxNumberOfTasks(tfs)
	minTasks := GetMinNumberOfTasks(tfs)

	return fmt.Sprintf("--------Task Completion Report--------\n"+
		"Directory: %s\n"+
		"Task Completion Percentage: %.2f\n"+
		"Mean Number of Tasks per File: %.2f +/- %.2f\n"+
		"Maximum number of tasks in one file: %d\n"+
		"Minimum number of tasks in one file: %d\n",
		path, percentage, mean, std, maxTasks, minTasks), nil
}

func GetPendingTasks(path string) ([]entities.TaskFile, error) {
	tf, err := getTaskFiles(path)
	if err != nil {
		return nil, err
	}

	return filterPendingTasks(tf), nil
}
