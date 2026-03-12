package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Contact struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func loadContacts() []Contact {
	var contacts []Contact
	data, err := os.ReadFile("contacts.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{}
		}

		fmt.Println("Error loading contacts:", err)
		return []Contact{}
	}

	err = json.Unmarshal(data, &contacts)
	if err != nil {
		fmt.Println("Error decoding json")
		return []Contact{}
	}

	return contacts

}

func saveContacts(contacts []Contact) {
	data, err := json.MarshalIndent(contacts, "", " ")
	if err != nil {
		fmt.Println("Could not encode into json data")
		return
	}

	err = os.WriteFile("contacts.json", data, 0644)
	if err != nil {
		fmt.Println("Error saving contacts:", err)
	}
}

func addContact(name, email, phone string) {
	contacts := loadContacts()

	var contact = Contact{
		Id:        len(contacts) + 1,
		Name:      name,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	contacts = append(contacts, contact)
	saveContacts(contacts)

	fmt.Printf("Contact added successfully (ID: %d)\n", contact.Id)
}

func listContacts() {
	contacts := loadContacts()

	if len(contacts) == 0 {
		fmt.Println("No contacts found")
		return
	}

	for _, contact := range contacts {
		fmt.Println(contact.Id, contact.Name, contact.Email, contact.Phone)
	}
}

func deleteContact(id int) {
	contacts := loadContacts()

	for i, contact := range contacts {
		if contact.Id == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			saveContacts(contacts)
			fmt.Printf("contact %d deleted successfully\n", id)
			return
		}
	}
	fmt.Printf("contact with ID %d not found\n", id)
}

func searchContact(query string) {
	contacts := loadContacts()
	found := false

	query = strings.ToLower(query)

	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Name), query) || strings.Contains(strings.ToLower(contact.Email), query) {
			fmt.Println(contact.Id, contact.Name, contact.Email, contact.Phone)
			found = true
		}
	}

	if !found {
		fmt.Println("No contacts found matching:", query)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: contact-book <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("usage: contact-book add <name> <email> <phone>")
			os.Exit(1)
		}
		addContact(os.Args[2], os.Args[3], os.Args[4])

	case "list":
		listContacts()

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage : contact-book delete <id>")
			os.Exit(1)
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID — must be a number")
			os.Exit(1)
		}

		deleteContact(id)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("usage : contact-book search <query>")
			os.Exit(1)
		}

		searchContact(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}

}
