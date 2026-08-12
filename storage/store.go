package storage

import "contactmanager/contact"

type Store interface {
	Load() ([]contact.Contact, error)
	Save([]contact.Contact) error
}
