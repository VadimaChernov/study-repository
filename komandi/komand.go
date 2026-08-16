package komandi

type Komand struct {
	tasks map[string]Task
}

func NewKomand() *Komand {
	komand := Komand{
		tasks: make(map[string]Task),
	}
	return &komand
}

func (k *Komand) Add(task Task) {
	k.tasks[task.Title] = task
}
func (k *Komand) List() map[string]Task {
	return k.tasks
}
func (k *Komand) Del(title string) string {
	_, ok := k.tasks[title]
	if !ok {
		return taskNotFound
	}
	delete(k.tasks, title)
	return ""
}
func (k *Komand) Done(title string) string {
	task, ok := k.tasks[title]
	if !ok {
		return taskNotFound
	}
	task.Done()
	k.tasks[title] = task
	return ""

}
