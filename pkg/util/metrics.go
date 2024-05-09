package util

import (
	"github.com/adolfo-ab/taskmd/pkg/util/entities"
	"math"
)

func GetTotalNumberOfTasks(tfs []entities.TaskFile) int {
	total := 0
	for _, tf := range tfs {
		total += len(tf.Tasks)
	}
	return total
}

func GetCompletionPercentage(tfs []entities.TaskFile) float64 {
	total := GetTotalNumberOfTasks(tfs)
	if total == 0 {
		return 0.0
	}

	completed := GetTotalNumberOfTasks(filterCompletedTasks(tfs))

	return float64(completed) / float64(total) * 100
}

func GetAverageNumberOfTasksPerFile(tfs []entities.TaskFile) float64 {
	numFiles := len(tfs)
	if numFiles == 0 {
		return 0.0
	}

	numTasks := 0
	for _, tf := range tfs {
		numTasks += len(tf.Tasks)
	}

	return float64(numTasks) / float64(numFiles)
}

func GetFilesWithMostTasks(tfs []entities.TaskFile) []entities.TaskFile {
	maxTfs := make([]entities.TaskFile, 0, len(tfs))
	maxNumTasks := len(tfs[0].Tasks)
	for _, tf := range tfs[1:] {
		if len(tf.Tasks) == maxNumTasks {
			maxTfs = append(maxTfs, tf)
		} else if len(tf.Tasks) > maxNumTasks {
			maxTfs = nil
			maxTfs = append(maxTfs, tf)
			maxNumTasks = len(tf.Tasks)
		}
	}
	return maxTfs
}

func GetMaxNumberOfTasks(tfs []entities.TaskFile) int {
	maxNumTasks := len(tfs[0].Tasks)
	for _, tf := range tfs[1:] {
		if len(tf.Tasks) > maxNumTasks {
			maxNumTasks = len(tf.Tasks)
		}
	}
	return maxNumTasks
}

func GetMinNumberOfTasks(tfs []entities.TaskFile) int {
	minNumTasks := len(tfs[0].Tasks)
	for _, tf := range tfs[1:] {
		if len(tf.Tasks) < minNumTasks {
			minNumTasks = len(tf.Tasks)
		}
	}
	return minNumTasks
}

func GetStd(tfs []entities.TaskFile, mean float64) float64 {
	var sum float64

	for _, tf := range tfs {
		deviation := float64(len(tf.Tasks)) - mean
		sum += deviation * deviation
	}

	variance := sum / float64(GetTotalNumberOfTasks(tfs))
	return math.Sqrt(variance)
}
