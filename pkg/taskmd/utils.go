package taskmd

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
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

func findNumberOfTasks(file string, completed bool) (int, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}
	contentStr := string(content)

	var tasks *regexp.Regexp
	if completed {
		tasks = completedTaskRegex
	} else {
		tasks = uncompletedTaskRegex
	}
	matches := tasks.FindAllStringIndex(contentStr, -1)
	return len(matches), nil
}

func findTaskCompletionPercentage(files []string) (float64, error) {
	completed := 0
	pending := 0

	for _, file := range files {
		c, err := findNumberOfTasks(file, true)
		if err != nil {
			return 0, err
		}
		completed += c

		p, err := findNumberOfTasks(file, false)
		if err != nil {
			return 0, err
		}
		pending += p
	}

	return float64(completed) / float64(completed+pending) * 100, nil
}
