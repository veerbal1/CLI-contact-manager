package contact

import (
	"errors"
	"strings"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var ErrContactNotFound = errors.New("contact not found")

func (c Contact) Validate() error {
	if c.Name == "" {
		return errors.New("Name is empty")
	} else if c.Phone == "" {
		return errors.New("Phone is empty")
	} else if c.Email == "" {
		return errors.New("Email is empty")
	} else if !strings.Contains(c.Email, "@") {
		return errors.New("Email must contain @")
	} else {
		return nil
	}
}

func (c *Contact) Update(name string, phone string, email string) error {
	contact := Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	err := contact.Validate()
	if err != nil {
		return err
	}
	c.Name = contact.Name
	c.Phone = contact.Phone
	c.Email = contact.Email
	return nil
}

func AddContact(contacts []Contact, contact Contact, nextID *int) ([]Contact, error) {
	err := contact.Validate()
	if err != nil {
		return contacts, err
	}
	contact.ID = *nextID
	*nextID++
	return append(contacts, contact), nil
}

func FindContactByID(contacts []Contact, id int) (*Contact, error) {
	for index := range contacts {
		contact := contacts[index]
		if contact.ID == id {
			return &contacts[index], nil
		}
	}
	return nil, ErrContactNotFound
}

func DeleteContactByID(contacts []Contact, id int) ([]Contact, error) {
	newContacts := make([]Contact, 0)
	found := false
	for index := range contacts {
		if contacts[index].ID != id {
			newContacts = append(newContacts, contacts[index])
		} else {
			found = true
			continue
		}
	}
	if !found {
		return contacts, ErrContactNotFound
	}
	return newContacts, nil
}

func NextID(contacts []Contact) int {
	maxID := 0

	for _, contact := range contacts {
		if contact.ID > maxID {
			maxID = contact.ID
		}
	}

	return maxID + 1
}
