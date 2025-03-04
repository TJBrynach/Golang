package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func createJSON() error {
	_, err := os.Stat("phonebook.json")
	if os.IsNotExist(err) {
		contacts := []Contact{}

		jsonData, err := json.MarshalIndent(contacts, "", "  ")
		if err != nil {
			return fmt.Errorf("error initialising json data: %v", err)
		}

		err = os.WriteFile("phonebook.json", jsonData, 0644)
		if err != nil {
			return fmt.Errorf("error creating the json file, %v", err)
		}
	}

	return nil

}

func appendJSON(contact Contact) error {
	file, err := os.ReadFile("phonebook.json")
	if err != nil {
		return fmt.Errorf("failed opening json to append")
	}

	// read (decode)
	var data []Contact
	err = json.Unmarshal(file, &data)
	if err != nil {
		return fmt.Errorf("failed to read and decode the file: %v", err)
	}

	// append
	data = append(data, contact)

	// write (encode)
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	err = os.WriteFile("phonebook.json", updatedJSON, 0644)
	if err != nil {
		return fmt.Errorf("error saving data: %v", err)
	}

	fmt.Println("successfully updated JSON")
	return nil

}

func loadContacts() ([]Contact, error) {
	file, err := os.Open("phonebook.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{}, nil
		}
		return nil, fmt.Errorf("error opening phonebook: %v", err)
	}
	defer file.Close()

	var contacts []Contact
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&contacts); err != nil {
		return nil, fmt.Errorf("error decoding phonebook json: %v", err)
	}

	return contacts, nil
}
