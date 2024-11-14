// TO-DO APP WITH CRUD OPERATION

package main

import (
	"fmt"
	"os"
)

type Todos struct {
	title     string
	completed bool
}

var todos []Todos

func showTodo() {
	if len(todos) == 0 {
		fmt.Println("No to-dos available.")
	} else {
		for i, todo := range todos {
			status := "Incomplete"
			if todo.completed {
				status = "Completed"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, todo.title, status)
		}
	}
}

func markTodoComp(index int) {
	if index < 0 || index >= len(todos) {
		fmt.Println("invalid index provided")
	}
	todos[index].completed = true
	fmt.Println("todo make as completed")
}

func updateTodo(index int, title string) {
	if index < 0 || index >= len(todos) {
		fmt.Println("invalid index provided")
	}
	todos[index].title = title
	fmt.Println("todo updated")
}

func deleteTodo(index int) {
	if index < 0 || index >= len(todos) {
		fmt.Println("invalid index")
	} else {
		todos = append(todos[:index], todos[index+1:]...)
		fmt.Println("todo deleted")
	}
}

func addTask(title string) {
	todo := Todos{title: title, completed: false}
	todos = append(todos, todo)
	println("todo added")
}

func main() {
	println("enter todo")
	var todo string
	var update string
	var index int
	var option int

	for {
		fmt.Println("1. Show To-dos")
		fmt.Println("2. Add To-do")
		fmt.Println("3. Update todo")
		fmt.Println("4. Mark a To-do as completed")
		fmt.Println("5. Delete a To-dos")
		fmt.Println("6, exit To-do App")

		fmt.Print("Input number for any of the options: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			showTodo()
		case 2:
			fmt.Println("enter a title for the To-do")
			fmt.Scan(&todo)
			addTask(todo)
			showTodo()

		case 3:
			fmt.Print("Enter the index of the to-do to update: ")
			fmt.Scan(&index)
			fmt.Print("Enter the new title for the to-do: ")
			fmt.Scan(&update)
			updateTodo(index-1, update)
			showTodo()
		case 4:
			fmt.Print("Enter the index of the to-do to mark as completed: ")
			fmt.Scan(&index)
			markTodoComp(index - 1)
			showTodo()
		case 5:
			deleteTodo(index - 1)
			showTodo()
		case 6:
			os.Exit(0)
		}
	}
}
