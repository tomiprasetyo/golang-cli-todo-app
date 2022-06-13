package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	listOfTodos := *t

	// check length of the index
	if index <= 0 || index > len(listOfTodos) {
		return errors.New("invalid index")
	}

	// if valid
	listOfTodos[index-1].CompletedAt = time.Now()
	listOfTodos[index-1].Done = true

	return nil
}

// method delete
func (t *Todos) Delete(index int) error {
	// get reference to the main data structure aka Todos
	listOfTodos := *t

	// check length of the index
	if index <= 0 || index > len(listOfTodos) {
		return errors.New("invalid index")
	}

	// dereference
	*t = append(listOfTodos[:index-1], listOfTodos[index:]...)

	return nil
}

// method load
func (t *Todos) Load(filename string) error {
	// read file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		// if file not found
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	// check the length of the file
	if len(file) == 0 {
		return err
	}

	// convert data from file to the struct object (unmarshal the data)
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// when add todo, write it back to the file
// method store
func (t *Todos) Store(filename string) error {
	// marshal the file
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	// write back data to the file
	return ioutil.WriteFile(filename, data, 0644)

}

// method print
func (t *Todos) Print() {
	// iterate the data
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}
