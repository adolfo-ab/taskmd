package util

import (
	"github.com/adolfo-ab/taskmd/pkg/util/entities"
	"testing"
)

func TestGetTotalNumberOfTasks(t *testing.T) {
	t1 := entities.NewTask("content1", "file1", true)
	t2 := entities.NewTask("content2", "file1", false)
	t3 := entities.NewTask("content3", "file2", true)
	t4 := entities.NewTask("content4", "file2", false)

	tf1 := entities.TaskFile{Path: "file1", Tasks: []entities.Task{t1, t2}}
	tf2 := entities.TaskFile{Path: "file2", Tasks: []entities.Task{t3, t4}}

	tf3 := entities.TaskFile{Path: "file3", Tasks: []entities.Task{}}

	var tests = []struct {
		name     string
		tfs      []entities.TaskFile
		expected int
	}{
		{name: "FourTasksTwoFiles", tfs: []entities.TaskFile{tf1, tf2}, expected: 4},
		{name: "ZeroTasksOneFile", tfs: []entities.TaskFile{tf3}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTotalNumberOfTasks(tt.tfs)
			if result != tt.expected {
				t.Errorf("Wrong number of tasks. Expected value: %d, actual value %d", tt.expected, result)
			}
		})
	}
}
