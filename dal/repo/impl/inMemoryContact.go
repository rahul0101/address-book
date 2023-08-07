package impl

import (
	"github.com/rahularcota/address-book/dal/model"
)

type InMemoryContactRepo struct {
	contacts []*model.Contact
}

func NewInMemoryContactRepo() *InMemoryContactRepo {
	return &InMemoryContactRepo{}
}

func (repo *InMemoryContactRepo) AddContact(contact *model.Contact) (int, error) {
	contact.Id = len(repo.contacts) + 1
	repo.contacts = append(repo.contacts, contact)
	return contact.Id, nil
}

func (repo *InMemoryContactRepo) GetContacts(ids []int) ([]*model.Contact, error) {
	var contacts []*model.Contact
	for _, id := range ids {
		contacts = append(contacts, repo.contacts[id-1])
	}
	return contacts, nil
}
