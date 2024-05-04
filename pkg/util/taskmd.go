package util

// TODO: Implement these 2 in metrics.go
/*func GetCompletionReport(path string) (string, error) {
}

func GetAverageNumberOfTasksPerFile(tasks[]) {
}*/

func GetCompletionPercentage(path string) (float64, error) {
	tf, err := getTaskFiles(path)
	if err != nil {
		return 0, err
	}
	total := GetTotalNumberOfTasks(tf)
	completed := GetTotalNumberOfTasks(filterCompletedTasks(tf))

	return float64(completed) / float64(total) * 100, nil
}

func GetPendingTasks(path string) ([]TaskFile, error) {
	tf, err := getTaskFiles(path)
	if err != nil {
		return nil, err
	}

	return filterPendingTasks(tf), nil
}
