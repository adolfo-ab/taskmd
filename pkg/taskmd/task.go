package taskmd

type Task struct {
	Content   string
	File      string
	Completed bool
}

func NewTask(content string, file string, completed bool) Task {
	return Task{
		Content:   content,
		File:      file,
		Completed: completed,
	}
}
