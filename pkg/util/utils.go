package util

import (
	"github.com/adolfo-ab/taskmd/pkg/util/entities"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func findMarkdownFiles(path string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(path, func(path string, di fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getTaskFiles(path string) ([]entities.TaskFile, error) {
	files, err := findMarkdownFiles(path)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	tfChan := make(chan entities.TaskFile)
	errChan := make(chan error)
	var taskFiles []entities.TaskFile

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			tf, err := entities.NewTaskFile(file)
			if err != nil {
				errChan <- err
				return
			}
			tfChan <- *tf
		}(file)
	}

	go func() {
		wg.Wait()
		close(tfChan)
		close(errChan)
	}()

	for tf := range tfChan {
		taskFiles = append(taskFiles, tf)
	}

	// Check if any errors occurred
	if len(errChan) > 0 {
		return nil, <-errChan // returns the first error encountered
	}

	return taskFiles, nil
}

func filterTasks(taskFiles []entities.TaskFile, condition func(entities.Task) bool) []entities.TaskFile {
	var filteredTaskFiles []entities.TaskFile
	for _, tf := range taskFiles {
		filteredTasks := make([]entities.Task, 0)
		for _, t := range tf.Tasks {
			if condition(t) {
				filteredTasks = append(filteredTasks, t)
			}
		}
		tf.Tasks = filteredTasks
		if len(filteredTasks) > 0 {
			filteredTaskFiles = append(filteredTaskFiles, tf)
		}
	}
	return filteredTaskFiles
}

func filterCompletedTasks(taskFiles []entities.TaskFile) []entities.TaskFile {
	return filterTasks(taskFiles, func(task entities.Task) bool {
		return task.Completed
	})
}

func filterPendingTasks(taskFiles []entities.TaskFile) []entities.TaskFile {
	return filterTasks(taskFiles, func(task entities.Task) bool {
		return !task.Completed
	})
}

func VerifyPathExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}
