package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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
	// // iterate the data
	// for i, item := range *t {
	// 	i++
	// 	fmt.Printf("%d - %s\n", i, item.Task)
	// }

	// add thirdparty library simple table
	// create instance of the table
	table := simpletable.New()

	// table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	// multidimensional array
	var cells [][]*simpletable.Cell

	// iterate every item in table
	for idx, item := range *t {
		idx++
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	// table body
	table.Body = &simpletable.Body{Cells: cells}

	// table footer
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "Your todos are here"},
	}}

	// table style
	table.SetStyle(simpletable.StyleUnicode)

	// print table
	table.Println()
}
