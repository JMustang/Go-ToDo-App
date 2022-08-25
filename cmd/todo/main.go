package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JMustang/Go-ToDo-App"
)

const (
	todofile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todofile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	default:
		fmt.Fprintln(os.Stdout, "Invalid command!")
		os.Exit(0)
	}
}
