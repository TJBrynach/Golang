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

	contact, err := searchContact(name)
	if err != nil {
		return fmt.Errorf("error searching for contact")
	}

	if name == contact.Name {
		return fmt.Errorf("user %v is already in your contact book", name)
	} else {
		fmt.Println("else")
		contact := Contact{
			Name:   name,
			Number: number,
		}
		appendJSON(contact)
	}
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
		return fmt.Errorf("error opening file to read: %v", name)
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

func searchContact(name string) (Contact, error) {
	// open file
	file, err := os.Open("phonebook.json")
	if err != nil {
		return Contact{}, fmt.Errorf("error opening file to read: %v", err)
	}

	// decode file
	var data []Contact
	decoder := json.NewDecoder(file)
	decoder.Decode(&data)

	// loop through and compare name
	found := false
	for _, contact := range data {
		if contact.Name == name {
			fmt.Println(contact.Display())
			found = true
			return contact, nil
		}

	}

	if !found {
		fmt.Println("No contact under that name")
	}

	// if so display
	// if not no user of that name
	return Contact{}, nil
}
