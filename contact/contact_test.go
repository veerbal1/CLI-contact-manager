package contact

import (
	"errors"
	"testing"
)

func TestContactValidate(t *testing.T) {
	tests := []struct {
		name    string
		contact Contact
		wantErr bool
	}{
		{
			name: "valid contact",
			contact: Contact{
				Name: "Veer", Phone: "123", Email: "veer@example.com",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			contact: Contact{
				Phone: "123", Email: "veer@example.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			contact: Contact{
				Name: "Veer", Phone: "123", Email: "veer",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		err := test.contact.Validate()

		if (err != nil) != test.wantErr {
			t.Errorf("%s: expected error=%v, got %v", test.name, test.wantErr, err)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.contact.Validate()

			if (err != nil) != test.wantErr {
				t.Errorf("expected error=%v, got %v", test.wantErr, err)
			}
		})
	}
}

func TestAddContact(t *testing.T) {
	contacts := []Contact{}
	id := NextID(contacts)
	if id != 1 {
		t.Errorf("Expected latest ID as: 1")
	}

	c1 := Contact{
		Name:  "veerb",
		Email: "veerbal@gmail.com",
		ID:    1,
		Phone: "123",
	}

	c2 := Contact{
		Name:  "mandeep",
		Email: "mandeep@gmail.com",
		ID:    2,
		Phone: "345",
	}

	if len(contacts) != 0 {
		t.Errorf("Expected 0 length")
	}

	newContacts, err := AddContact(contacts, c1, &id)
	if err != nil {
		t.Errorf("Failed to add new contact: 1")
	}
	contacts = newContacts
	if contacts[0].ID != 1 {
		t.Errorf("Expected contact ID as: 1")
	}

	if len(contacts) != 1 {
		t.Errorf("Expected 1 length")
	}

	newContacts2, err := AddContact(contacts, c2, &id)
	if err != nil {
		t.Errorf("Failed to add new contact: 1")
	}
	contacts = newContacts2

	if len(contacts) != 2 {
		t.Errorf("Expected 2 length")
	}

	if contacts[1].ID != 2 {
		t.Errorf("Expected contact ID as: 2")
	}

	if id != 3 {
		t.Errorf("Expected latest ID as: 3")
	}
}

func TestFindContactByID(t *testing.T) {
	contacts := []Contact{{
		Name:  "veerb",
		Email: "veerbal@gmail.com",
		ID:    1,
		Phone: "123",
	}, {
		Name:  "mandeep",
		Email: "mandeep@gmail.com",
		ID:    2,
		Phone: "345",
	}}
	c, err := FindContactByID(contacts, 2)
	if err != nil {
		t.Errorf("Not expecting an error to find 2nd element")
		return
	}

	if c.ID != 2 {
		t.Errorf("Expecting contact ID as 2")
	}

	if c.Name != "mandeep" {
		t.Errorf("Expecting name as 'mandeep'")
	}
}

func TestFindContactByIDNotFound(t *testing.T) {
	contacts := []Contact{{
		Name:  "veerb",
		Email: "veerbal@gmail.com",
		ID:    1,
		Phone: "123",
	}, {
		Name:  "mandeep",
		Email: "mandeep@gmail.com",
		ID:    2,
		Phone: "345",
	}}
	_, err := FindContactByID(contacts, 3)
	if !errors.Is(err, ErrContactNotFound) {
		t.Errorf("expecting err: ErrContactNotFound")
	}
}
