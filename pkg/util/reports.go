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

	color := RedColor
	if percentage >= 50.0 {
		color = GreenColor
	}

	return fmt.Sprintf("--------Task Completion Report--------\n"+
		"Directory: %s\n"+
		"Task completion percentage: %s%.2f%%%s\n"+
		"Mean number of tasks/fileile: %.2f +/- %.2f\n"+
		"Max number of tasks/file: %d\n"+
		"Min number of tasks/file: %d\n",
		path, color, percentage, ResetColor, mean, std, maxTasks, minTasks), nil
}

func GetPendingTasks(path string) ([]entities.TaskFile, error) {
	tf, err := getTaskFiles(path)
	if err != nil {
		return nil, err
	}

	return filterPendingTasks(tf), nil
}
