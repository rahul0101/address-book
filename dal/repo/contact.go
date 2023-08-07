package repo

import "github.com/rahularcota/address-book/dal/model"

type IContactRepo interface {
	AddContact(contact *model.Contact) (int, error)
	GetContacts(ids []int) ([]*model.Contact, error)
}
