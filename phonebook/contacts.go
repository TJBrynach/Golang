package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func (c Contact) Display() string {
	return fmt.Sprintf("%s: %s", c.Name, c.Number)
}

func createContact(name string, number string) error {
	contact := Contact{
		Name:   name,
		Number: number,
	}

	appendJSON(contact)

	return nil
}

func listContacts() error {
	// open file
	file, err := os.Open("phonebook.json")
	if err != nil {
		return fmt.Errorf("error opening phonebook: %v", err)
	}

	// decode file
	var data []Contact
	encoder := json.NewDecoder(file)
	encoder.Decode(&data)

	for _, contact := range data {
		fmt.Println(contact.Display())
	}

	return nil
}

func deleteContact(name string) error {
	// open file
	file, err := os.Open("phonebook.json")
	if err != nil {
		return fmt.Errorf("error opening file to read: %v", err)
	}

	// decode file into an object
	var jsonData []Contact
	decoder := json.NewDecoder(file)
	decoder.Decode(&jsonData)

	// loop through file and write the data into a new file if name != name
	var newList []Contact
	for _, contact := range jsonData {
		if contact.Name != name {
			newList = append(newList, contact)
		}
	}
	// open file once more
	file, err = os.OpenFile("phonebook.json", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file to write: %v", err)
	}

	// encode file into file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(newList)

	return nil
}
