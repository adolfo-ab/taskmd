package taskmd

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

func findTasksInFile(file string) ([]Task, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trimming leading and trailing white spaces
		if strings.HasPrefix(line, "- [ ] ") {
			tasks = append(tasks, NewTask(strings.TrimSpace(line[5:]), false))
		} else if strings.HasPrefix(line, "- [x] ") {
			tasks = append(tasks, NewTask(strings.TrimSpace(line[5:]), true))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func findTasksInFiles(files []string) ([]Task, error) {
	var wg sync.WaitGroup
	taskChan := make(chan []Task)
	errChan := make(chan error)
	var tasks []Task

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			t, err := findTasksInFile(file)
			if err != nil {
				errChan <- err
				return
			}
			taskChan <- t
		}(file)
	}

	go func() {
		wg.Wait()
		close(taskChan)
		close(errChan)
	}()

	for t := range taskChan {
		tasks = append(tasks, t...)
	}

	// Check if any errors occurred
	if len(errChan) > 0 {
		return nil, <-errChan // returns the first error encountered
	}

	return tasks, nil
}

func filterTasks(tasks []Task, condition func(Task) bool) []Task {
	filtered := make([]Task, 0)
	for _, task := range tasks {
		if condition(task) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func filterCompletedTasks(tasks []Task) []Task {
	return filterTasks(tasks, func(task Task) bool {
		return task.Completed
	})
}

func filterPendingTasks(tasks []Task) []Task {
	return filterTasks(tasks, func(task Task) bool {
		return !task.Completed
	})
}

func VerifyPathExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}
