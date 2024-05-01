package taskmd

type Task struct {
	Content   string
	Completed bool
}

func NewTask(content string, completed bool) Task {
	return Task{
		Content:   content,
		Completed: completed,
	}
}
