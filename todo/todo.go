package main

import "fmt"

type Todo struct {
	id        int
	todo      string
	completed bool
}

type Todos []Todo

func (t *Todos) createTodo(id int, todo string, completed bool) Todo {
	newTodo := Todo{id: id, todo: todo, completed: completed}
	*t = append(*t, newTodo)
	return newTodo
}

func (t *Todos) getAllTodos() {
	for _, todo := range *t {
		fmt.Println("TODO ID : ", todo.id, "\n", " TODO : ", todo.todo, "\n", " TODO STAUS : ", todo.completed)
		fmt.Println("**********************")
	}

}

func (t *Todos) deleteTodo(id int) {
	for i, v := range *t {
		if v.id == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			break
		}
	}
}


