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
	tf4 := entities.TaskFile{Path: "file4"}

	var tests = []struct {
		name     string
		tfs      []entities.TaskFile
		expected int
	}{
		{name: "FourTasksTwoFiles", tfs: []entities.TaskFile{tf1, tf2}, expected: 4},
		{name: "ZeroTasksOneFile", tfs: []entities.TaskFile{tf3}, expected: 0},
		{name: "NoTasksAtAll", tfs: []entities.TaskFile{tf4}, expected: 0},
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

func TestGetCompletionPercentage(t *testing.T) {
	t1 := entities.NewTask("content1", "file1", true)
	t2 := entities.NewTask("content2", "file1", false)

	t3 := entities.NewTask("content3", "file2", true)
	t4 := entities.NewTask("content4", "file2", true)

	t5 := entities.NewTask("content5", "file3", false)
	t6 := entities.NewTask("content6", "file3", false)

	tf1 := entities.TaskFile{Path: "file1", Tasks: []entities.Task{t1, t2}}
	tf2 := entities.TaskFile{Path: "file2", Tasks: []entities.Task{t3, t4}}
	tf3 := entities.TaskFile{Path: "file3", Tasks: []entities.Task{t5, t6}}
	tf4 := entities.TaskFile{Path: "file4", Tasks: []entities.Task{}}
	tf5 := entities.TaskFile{Path: "file5"}

	var tests = []struct {
		name     string
		tfs      []entities.TaskFile
		expected float64
	}{
		{name: "HalfCompleted", tfs: []entities.TaskFile{tf1}, expected: 50.0},
		{name: "HalfCompletedThreeTaskFiles", tfs: []entities.TaskFile{tf1, tf2, tf3}, expected: 50.0},
		{name: "AllCompleted", tfs: []entities.TaskFile{tf2}, expected: 100.0},
		{name: "NoneCompleted", tfs: []entities.TaskFile{tf3}, expected: 0},
		{name: "EmptyTaskSlice", tfs: []entities.TaskFile{tf4}, expected: 0},
		{name: "NoTasksAtAll", tfs: []entities.TaskFile{tf5}, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCompletionPercentage(tt.tfs)
			if result != tt.expected {
				t.Errorf("Incorrect completion percentage. Expected value: %.2f, actual value %.2f", tt.expected, result)
			}
		})
	}
}

func TestGetAverageNumberOfTasksPerFile(t *testing.T) {
	t1 := entities.NewTask("content1", "file1", true)
	t2 := entities.NewTask("content2", "file1", false)
	t3 := entities.NewTask("content3", "file2", true)

	tf1 := entities.TaskFile{Path: "file1", Tasks: []entities.Task{t1, t2}}
	tf2 := entities.TaskFile{Path: "file2", Tasks: []entities.Task{t3}}

	t5 := entities.NewTask("content1", "file3", true)
	t6 := entities.NewTask("content2", "file4", false)
	t7 := entities.NewTask("content3", "file5", true)
	t8 := entities.NewTask("content4", "file6", true)

	tf3 := entities.TaskFile{Path: "file3", Tasks: []entities.Task{t5}}
	tf4 := entities.TaskFile{Path: "file4", Tasks: []entities.Task{t6}}
	tf5 := entities.TaskFile{Path: "file5", Tasks: []entities.Task{t7}}
	tf6 := entities.TaskFile{Path: "file6", Tasks: []entities.Task{t8}}

	tf7 := entities.TaskFile{Path: "file7", Tasks: []entities.Task{}}
	tf8 := entities.TaskFile{Path: "file8", Tasks: []entities.Task{}}

	tf9 := entities.TaskFile{Path: "file9"}
	tf10 := entities.TaskFile{Path: "file10"}

	var tests = []struct {
		name     string
		tfs      []entities.TaskFile
		expected float64
	}{
		{name: "DifferentNumberOfTasks", tfs: []entities.TaskFile{tf1, tf2}, expected: 1.5},
		{name: "OneTaskPerFile", tfs: []entities.TaskFile{tf3, tf4, tf5, tf6}, expected: 1},
		{name: "EmptyTaskSlice", tfs: []entities.TaskFile{tf7, tf8}, expected: 0.0},
		{name: "NoTasksAtAll", tfs: []entities.TaskFile{tf9, tf10}, expected: 0.0},
		{name: "NoTaskFiles", tfs: []entities.TaskFile{}, expected: 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAverageNumberOfTasksPerFile(tt.tfs)
			if result != tt.expected {
				t.Errorf("Incorrect average. Expected value: %.2f, actual value %.2f", tt.expected, result)
			}
		})
	}

}
