package komandi

import "time"

type Task struct {
	Title    string
	Text     string
	IsDone   bool
	CreateAt time.Time
	DoneAt   *time.Time
}

func NewTask(title string, text string) Task {
	return Task{
		Title:    title,
		Text:     text,
		IsDone:   false,
		CreateAt: time.Now(),
		DoneAt:   nil,
	}
}

func (t *Task) Done() {
	t.IsDone = true
	now := time.Now()
	t.DoneAt = &now
}
