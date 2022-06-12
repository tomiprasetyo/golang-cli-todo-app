package todo

import (
	"errors"
	"time"
)

// create data strcuture
type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// collection of items
type Todos []Item

// method add
func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	// dereference the list
	*t = append(*t, todo)
}

// method completed
func (t *Todos) Complete(index int) error {
	// get reference to the main data structure aka Todos
	listTodos := *t

	// check length of the index
	if index <= 0 || index > len(listTodos) {
		return errors.New("invalid index")
	}

	// if valid
	listTodos[index-1].CompletedAt = time.Now()
	listTodos[index-1].Done = true

	return nil
}
