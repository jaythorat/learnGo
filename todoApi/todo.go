package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func (t *Todos) createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

	*t = append(*t, todo)
	// Now 'note' contains the posted data
    fmt.Fprintf(w, "Received note: %+v\n", todo)
}

func (t *Todos) getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := make(map[int]Todo)

	for i, todo := range *t {
		todos[i] = todo
	}
	fmt.Println(todos)

	json.NewEncoder(w).Encode(todos)

}
