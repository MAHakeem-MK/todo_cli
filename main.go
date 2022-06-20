package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Color string
type Symbol string

const (
	RESET Color = "\033[0m"
	RED   Color = "\033[31m"
	GREEN Color = "\033[32m"
)

const (
	TODO Symbol = "\u2717"
	DONE Symbol = "\u2713"
)

type TodoItem struct {
	title  string
	isDone bool
}

func main() {
	RenderTodos(ReadTodos())
}

func ReadTodos() []TodoItem {
	var todos []TodoItem
	f, err := os.Open("./todolist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		stateChar := text[0]
		if stateChar == '-' {
			todos = append(todos, TodoItem{title: text[1:], isDone: false})
		}
		if stateChar == '+' {
			todos = append(todos, TodoItem{title: text[1:], isDone: true})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return todos
}

func WriteTodos() {

}

func RenderTodos(todos []TodoItem) {
	var todo []TodoItem
	var done []TodoItem
	for _, v := range todos {
		if v.isDone {
			done = append(done, v)
		} else {
			todo = append(todo, v)
		}
	}
	for _, v := range todo {
		fmt.Println(RED, TODO, v.title, RESET)
	}
	for _, v := range done {
		fmt.Println(GREEN, DONE, v.title, RESET)
	}
}