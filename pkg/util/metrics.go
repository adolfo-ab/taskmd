package util

import "github.com/adolfo-ab/taskmd/pkg/util/entities"

func GetTotalNumberOfTasks(tfs []entities.TaskFile) int {
	total := 0
	for _, tf := range tfs {
		total += len(tf.Tasks)
	}
	return total
}

func GetCompletionPercentage(tfs []entities.TaskFile) float64 {
	total := GetTotalNumberOfTasks(tfs)
	completed := GetTotalNumberOfTasks(filterCompletedTasks(tfs))

	return float64(completed) / float64(total) * 100
}

func GetAverageNumberOfTasksPerFile(tfs []entities.TaskFile) float64 {
	numFiles := len(tfs)

	numTasks := 0
	for _, tf := range tfs {
		numTasks += len(tf.Tasks)
	}

	return float64(numTasks / numFiles)
}
