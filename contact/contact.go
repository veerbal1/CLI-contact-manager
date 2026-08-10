package contact

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func (c *Contact) Update(name string, phone string, email string) {
	c.Email = email
	c.Phone = phone
	c.Name = name
}

func AddContact(contacts []Contact, contact Contact, nextID *int) []Contact {
	contact.ID = *nextID
	*nextID++
	return append(contacts, contact)
}

func FindContactByID(contacts []Contact, id int) (*Contact, bool) {
	for index := range contacts {
		contact := contacts[index]
		if contact.ID == id {
			return &contacts[index], true
		}
	}
	return nil, false
}

func DeleteContactByID(contacts []Contact, id int) ([]Contact, bool) {
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
	return newContacts, found
}
