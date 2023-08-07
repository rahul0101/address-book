package service

import (
	"github.com/rahularcota/address-book/constants"
	"github.com/rahularcota/address-book/dtos"
)

type IContactService interface {
	AddContact(contact *dtos.Contact) error
	SearchContacts(searchType constants.IndexType, searchString string) ([]*dtos.Contact, error)
}
