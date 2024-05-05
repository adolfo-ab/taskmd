package entities

const (
	Completed = "- [x] "
	Pending   = "- [ ] "
)

type Task struct {
	Content   string
	Completed bool
}

func NewTask(content string, file string, completed bool) Task {
	return Task{
		Content:   content,
		Completed: completed,
	}
}
