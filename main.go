package main

import (
	"bufio"
	"contactmanager/contact"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeUserInput(scanner *bufio.Scanner, title string) (string, error) {
	fmt.Print(title)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("input ended")
	}
	userInput := scanner.Text()
	userInput = strings.TrimSpace(userInput)
	return userInput, nil
}

func listContacts(contacts []contact.Contact) {
	if len(contacts) == 0 {
		fmt.Println("Empty contact list")
		return
	}
	for _, contact := range contacts {
		fmt.Println(contact.ID, contact.Name, contact.Phone, contact.Email)
	}
}

func contactNamesByID(contacts []contact.Contact) map[int]string {
	contactsMap := make(map[int]string)
	for index := range contacts {
		contactsMap[contacts[index].ID] = contacts[index].Name
	}
	return contactsMap
}

func main() {
	if len(os.Args) > 1 {
		helpArg := os.Args[1]
		if helpArg == "help" {
			fmt.Println("Usage: contactmanager help")
			return
		}
		fmt.Println("Invalid argument")
		return
	}
	contacts := make([]contact.Contact, 0)
	nextID := 1

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("---------------------")
		fmt.Println("Add Contact: Press 1")
		fmt.Println("List contacts: Press 2")
		fmt.Println("Find contact by ID: Press 3")
		fmt.Println("Update contact by ID: Press 4")
		fmt.Println("Delete contact by ID: Press 5")
		fmt.Println("Exit: Press 6")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Input error:", err)
			}
			return
		}
		userInput := scanner.Text()
		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "1":
			name, err := takeUserInput(scanner, "Enter name: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			phone, err := takeUserInput(scanner, "Enter phone: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			email, err := takeUserInput(scanner, "Enter email: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			contactNew := contact.Contact{
				Name:  name,
				Phone: phone,
				Email: email,
			}

			contacts, err = contact.AddContact(contacts, contactNew, &nextID)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "2":
			listContacts(contacts)
		case "3":
			id, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			textID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			contact_, err := contact.FindContactByID(contacts, textID)
			if errors.Is(err, contact.ErrContactNotFound) {
				fmt.Println("Could not find contact with ID:", textID)
				continue
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println((*contact_).ID, contact_.Name, contact_.Phone, contact_.Email)
			continue
		case "4":
			// ask user for id
			userInput, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			textID, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			}
			contact_, err := contact.FindContactByID(contacts, textID)
			if errors.Is(err, contact.ErrContactNotFound) {
				fmt.Println("Could not find contact with ID:", textID)
				continue
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(contact_.ID, contact_.Name, contact_.Phone, contact_.Email)
			// Ask for the new name, phone, and email.
			name, err := takeUserInput(scanner, "Enter name: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			phone, err := takeUserInput(scanner, "Enter phone: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			email, err := takeUserInput(scanner, "Enter email: ")
			if err != nil {
				fmt.Println(err)
				return
			}

			err = contact_.Update(name, phone, email)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Contact updated successfully")
		case "5":
			// ask user for id
			userInput, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println(err)
				return
			}
			textID, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// delete contact by id
			updatedContacts, err := contact.DeleteContactByID(contacts, textID)
			if errors.Is(err, contact.ErrContactNotFound) {
				fmt.Println("Could not find contact with ID:", textID)
				continue
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			contacts = updatedContacts
			fmt.Println("Contact deleted successfully")
			continue
		case "6":
			return
		default:
			fmt.Println("Invalid argument")
			continue
		}
	}
}
