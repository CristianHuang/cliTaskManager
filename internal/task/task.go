package task

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type List []Task

func (l *List) Add(title string) {
	id := len(*l) + 1

	task := Task{
		ID:    id,
		Title: title,
		Done:  false,
	}

	*l = append(*l, task)
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
