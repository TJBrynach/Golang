package main

import (
	"fmt"
	"os"
)

// this is where we will illustrate our cli client

func HandleCLI() {
	fmt.Println("Debug: CLI started:")
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./todolist [command] [arguments]")
		fmt.Println("Available commands:")
		return
	}
	fmt.Println(os.Args)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./todolist add \"Task title\"")
			return
		}
		title := os.Args[2]
		err := createTask(title, fileName)

		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		// fmt.Println("Task Added: ", task.Title)

	case "list":
		if len(os.Args) < 2 {
			fmt.Println("Usage: . list")
			return
		}
		// tasks, err := listTasks(fileName)
		// if err != nil {
		// 	fmt.Println("Error loading tasks:", err)
		// 	return
		// }

		// for _, task := range tasks {
		// 	fmt.Println(task.Display())
		// }
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: . delete \"title\"")
		}

		title := os.Args[2]
		err := deleteTasks(title, fileName)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("task deleted: ", title)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Incorrect usage")
		}
		title := os.Args[2]
		err := completeTask(title, fileName)

		if err != nil {
			fmt.Println("error", err)
		}
	default:
		fmt.Println("Unknown command: ", command)
		fmt.Println("Available commands are: add, list, delete, complete")
	}
}
