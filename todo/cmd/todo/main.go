package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomiprasetyo/golang-cli-todo-app/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	// make flag
	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	// unite
	todos := &todo.Todos{}

	// load todo when start the application
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// switch case
	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
