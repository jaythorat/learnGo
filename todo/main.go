package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var input int
	var todos Todos
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter choice : ")
		fmt.Println("1. Get all TODO  ")
		fmt.Println("2. Create TODO ")
		fmt.Println("3. Delete TODO")
		fmt.Scanln(&input)

		choice := input

		switch choice {
		case 1:
			todos.getAllTodos()

		case 2:
			var id int
			var todo string
			fmt.Println("Enter todo Id :")
			fmt.Scanln(&id)
			fmt.Println("Enter todo TEXT :")
			todo,err := reader.ReadString('\n')
			todo = strings.TrimSpace(todo)
			fmt.Println(todo,"TODO IS HERE",len(todo),err)
			if todo != "" {
				todos.createTodo(id, todo, false)
				fmt.Println("todo created Successfully")
			} else {
				fmt.Println("Empty todo, not added.")
			}

		case 3:
			fmt.Println("Enter ID Of TODO to delete : ")
			var id int
			fmt.Scanln(&id)
			todos.deleteTodo(id)

		default:
			fmt.Println("Exiting program.")
			return
		}

	}

}
