package taskmd

import (
	"regexp"
)

var (
	completedTaskRegex   = regexp.MustCompile(`\[x]`)
	uncompletedTaskRegex = regexp.MustCompile(`\[ ]`)
)

func GetTaskCompletionPercentage(path string) (float64, error) {
	files, err := findMarkdownFiles(path)
	percentage, err := findTaskCompletionPercentage(files)

	return percentage, err
}
