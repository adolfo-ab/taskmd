package entities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TaskFile struct {
	Path  string
	Tasks []Task
}

func NewTaskFile(name string) (*TaskFile, error) {
	var tf TaskFile
	tf.Path = name
	err := tf.parse()
	if err != nil {
		return nil, err
	}
	return &tf, nil
}

func (tf *TaskFile) GetTasks() []Task {
	return tf.Tasks
}

func (tf *TaskFile) DeleteTask(index int) error {
	if index < 0 || index >= len(tf.Tasks) {
		return fmt.Errorf("invalid index: %d", index)
	}

	tf.Tasks = append(tf.Tasks[:index], tf.Tasks[index+1:]...)

	return nil
}

func (tf *TaskFile) parse() error {
	dat, err := os.ReadFile(tf.Path)
	if err != nil {
		return err
	}

	var tasks []Task
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trimming leading and trailing white spaces
		if strings.HasPrefix(line, Pending) {
			tasks = append(tasks, NewTask(strings.TrimSpace(line[5:]), tf.Path, false))
		} else if strings.HasPrefix(line, Completed) {
			tasks = append(tasks, NewTask(strings.TrimSpace(line[5:]), tf.Path, true))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	tf.Tasks = tasks
	return nil
}
