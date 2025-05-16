package main

import (
	"net/http"
)

func main() {
	var todos Todos

	todos = append(todos, Todo{10,"Jay Thorat",false})

	http.HandleFunc("/getAllNotes", todos.getAllTodos)
	http.HandleFunc("/createtodo", todos.createTodo)
	
	http.ListenAndServe(":8080", nil)

}
