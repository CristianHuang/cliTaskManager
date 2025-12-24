package task

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	SubTasks  []SubTask `json:"sub_tasks"`
}
type SubTask struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type List []Task

func (l *List) Add(title string) {
	id := len(*l) + 1

	task := Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		SubTasks:  []SubTask{},
	}

	*l = append(*l, task)
}

func (l *List) AddSubTask(taskID int, title string) {
	for i, t := range *l {
		if t.ID == taskID {
			(*l)[i].SubTasks = append((*l)[i].SubTasks, SubTask{
				ID:        len((*l)[i].SubTasks) + 1,
				Title:     title,
				Done:      false,
				CreatedAt: time.Now(),
			})
			return
		}
	}
}

func (l *List) Complete(id int) bool {
	for i, t := range *l {
		if t.ID == id {
			(*l)[i].Done = true
			return true
		}
	}
	return false
}

func (l *List) Delete(id int) bool {
	for i, t := range *l {
		if t.ID == id {
			*l = append((*l)[:i], (*l)[i+1:]...)
			return true
		}
	}
	return false
}

func (l *List) Edit(id int, title string) bool {
	for i, t := range *l {
		if t.ID == id {
			(*l)[i].Title = title
			return true
		}
	}
	return false
}
