package main

import (
	"fmt"
	"os"
)

func HandleCLI() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("incorrect statement")
			return
		}
		name := os.Args[2]
		number := os.Args[3]
		err := createContact(name, number)
		if err != nil {
			fmt.Println("Error occured creating contact:", err)
			return
		}

		fmt.Println("successfully added Contact:", name)
	case "list":
		if len(os.Args) < 2 {
			fmt.Println("incorrect number of arguments")
			return
		}

		err := listContacts()
		if err != nil {
			fmt.Println("error listing contacts")
			return
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments to delete")
			return
		}

		name := os.Args[2]
		err := deleteContact(name)
		if err != nil {
			fmt.Println("error deleting user", err)
			return
		}

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("incorrect number of arguments")
			return
		}

		name := os.Args[2]

		_, err := searchContact(name)
		if err != nil {
			fmt.Println("error searching ", err)
		}

	default:
		fmt.Println("inaccurate command the options are: add, list")
	}

}
