package taskmd

func GetCompletionPercentage(path string) (float64, error) {
	files, err := findMarkdownFiles(path)
	if err != nil {
		return 0, err
	}

	tasks, err := findTasksInFiles(files)
	if err != nil {
		return 0, err
	}
	completed := filterCompletedTasks(tasks)

	return float64(len(completed)) / float64(len(tasks)) * 100, nil
}

func GetPendingTasks(path string) ([]Task, error) {
	files, err := findMarkdownFiles(path)
	if err != nil {
		return nil, err
	}

	tasks, err := findTasksInFiles(files)
	if err != nil {
		return nil, err
	}

	return filterPendingTasks(tasks), nil
}
