package storage

import (
	"contactmanager/contact"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type JSONStore struct {
	fileName string
}

type MemoryStore struct {
	contacts []contact.Contact
}

func NewMemoryStore(contacts []contact.Contact) MemoryStore {
	return MemoryStore{
		contacts,
	}
}

func NewJSONStore(fileName string) JSONStore {
	return JSONStore{
		fileName: fileName,
	}
}

func (s JSONStore) Load() ([]contact.Contact, error) {
	file, err := os.Open(s.fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []contact.Contact{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var contacts []contact.Contact

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&contacts)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return []contact.Contact{}, nil
		}

		return nil, err
	}
	return contacts, nil
}

func (s JSONStore) Save(contacts []contact.Contact) error {
	file, err := os.Create(s.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

func (s MemoryStore) Load() ([]contact.Contact, error) {
	return s.contacts, nil
}

func (s *MemoryStore) Save(contacts []contact.Contact) error {
	s.contacts = contacts
	return nil
}
