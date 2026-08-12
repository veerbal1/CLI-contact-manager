package storage

import (
	"contactmanager/contact"
	"path/filepath"
	"testing"
)

func TestJSONStoreSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	var fileName string = "contacts.json"
	path := filepath.Join(dir, fileName)
	store := NewJSONStore(path)

	contacts := []contact.Contact{{
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

	err := store.Save(contacts)
	if err != nil {
		t.Errorf("unexpected save error: %v", err)
		return
	}

	contactsLoaded, err := store.Load()
	if err != nil {
		t.Errorf("unexpected loading error: %v", err)
		return
	}

	if len(contactsLoaded) != len(contacts) {
		t.Errorf("loaded contacts length mismatch: expected length: %v, got: %v", len(contacts), len(contactsLoaded))
		return
	}

	if contactsLoaded[0].Name != contacts[0].Name {
		t.Errorf("loaded contacts name mismatch: expected: %v, got: %v", contacts[0].Name, contactsLoaded[0].Name)
	}

	if contactsLoaded[0].ID != contacts[0].ID {
		t.Errorf("loaded contacts ID mismatch: expected: %v, got: %v", contacts[0].ID, contactsLoaded[0].ID)
	}

	if contactsLoaded[1].Name != contacts[1].Name {
		t.Errorf("loaded contacts name mismatch: expected: %v, got: %v", contacts[1].Name, contactsLoaded[1].Name)
	}

	if contactsLoaded[1].ID != contacts[1].ID {
		t.Errorf("loaded contacts ID mismatch: expected: %v, got: %v", contacts[1].ID, contactsLoaded[1].ID)
	}
}
