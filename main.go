package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addContact(contacts []string, contact string) []string {
	return append(contacts, contact)
}

func listContacts(contacts []string) {
	if len(contacts) == 0 {
		fmt.Println("Empty contact list")
		return
	}
	for index, contact := range contacts {
		fmt.Println(index+1, contact)
	}
}

func main() {
	contacts := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("---------------------")
		fmt.Println("Add Contact: Press 1")
		fmt.Println("List contacts: Press 2")
		fmt.Println("Exit: Press 3")
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
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()
			name = strings.TrimSpace(name)
			if name == "" {
				fmt.Println("Provided empty name")
				continue
			}

			contacts = addContact(contacts, name)
		case "2":
			listContacts(contacts)
		case "3":
			return
		default:
			fmt.Println("Invalid argument")
			continue
		}
	}
}
