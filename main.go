package main

import (
	"bufio"
	"contactmanager/contact"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addContact(contacts []contact.Contact, contact contact.Contact, nextID *int) []contact.Contact {
	contact.ID = *nextID
	*nextID++
	return append(contacts, contact)
}

func takeUserInput(scanner *bufio.Scanner, title string) (string, error) {
	fmt.Print(title)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Input error:", err)
		}
		return "", fmt.Errorf("got error")
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
				fmt.Println("Something went wrong")
				return
			}
			if name == "" {
				fmt.Println("Provided empty name")
				continue
			}

			phone, err := takeUserInput(scanner, "Enter phone: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			if phone == "" {
				fmt.Println("Provided empty phone number")
				continue
			}

			email, err := takeUserInput(scanner, "Enter email: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			if email == "" {
				fmt.Println("Provided email is wrong")
				continue
			}

			contactNew := contact.Contact{
				Name:  name,
				Phone: phone,
				Email: email,
			}

			contacts = contact.AddContact(contacts, contactNew, &nextID)
		case "2":
			listContacts(contacts)
		case "3":
			id, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			textID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Something went wrong")
				continue
			}
			contact, found := contact.FindContactByID(contacts, textID)
			if !found {
				fmt.Println("Could not find contact with ID: ", textID)
				continue
			}
			fmt.Println((*contact).ID, contact.Name, contact.Phone, contact.Email)
			continue
		case "4":
			// ask user for id
			userInput, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			textID, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Something went wrong")
				continue
			}
			contact, found := contact.FindContactByID(contacts, textID)
			if !found {
				fmt.Println("Could not find contact with ID: ", textID)
				continue
			}
			fmt.Println((*contact).ID, contact.Name, contact.Phone, contact.Email)
			// Ask for the new name, phone, and email.
			name, err := takeUserInput(scanner, "Enter name: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			phone, err := takeUserInput(scanner, "Enter phone: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			email, err := takeUserInput(scanner, "Enter email: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			// validate name, phone, email
			if name == "" {
				fmt.Println("Provided empty name")
				continue
			}
			if phone == "" {
				fmt.Println("Provided empty phone number")
				continue
			}
			if email == "" {
				fmt.Println("Provided email is wrong")
				continue
			}

			contact.Update(name, phone, email)
			fmt.Println("Contact updated successfully")
		case "5":
			// ask user for id
			userInput, err := takeUserInput(scanner, "Type ID: ")
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			textID, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Something went wrong")
				continue
			}
			// delete contact by id
			updatedContacts, found := contact.DeleteContactByID(contacts, textID)
			if !found {
				fmt.Println("Could not find contact with ID: ", textID)
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
