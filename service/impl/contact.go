package impl

import (
	"fmt"
	"github.com/rahularcota/address-book/constants"
	"github.com/rahularcota/address-book/dal/repo"
	"github.com/rahularcota/address-book/dtos"
	"github.com/rahularcota/address-book/managers"
	"github.com/rahularcota/address-book/mappers"
	"strings"
	"time"
)

type ContactService struct {
	contactRepo  repo.IContactRepo
	indexManager managers.IIndexManager
}

func NewContactService(contactRepo repo.IContactRepo, indexManager managers.IIndexManager) *ContactService {
	contactService := &ContactService{
		contactRepo:  contactRepo,
		indexManager: indexManager,
	}

	//for i := 1; i < 10000000; i++ {
	//	contactService.AddContact(&dtos.Contact{PhoneNumber: strconv.Itoa(i)})
	//}
	return contactService
}

func (svc *ContactService) AddContact(contact *dtos.Contact) error {
	idx, err := svc.contactRepo.AddContact(mappers.MapContactDTOToModel(contact))
	if err != nil {
		return err
	}
	addToIndexReq := &dtos.AddToIndexRequest{
		Id:          idx,
		FirstName:   strings.ToLower(contact.FirstName),
		LastName:    strings.ToLower(contact.LastName),
		PhoneNumber: strings.ToLower(contact.PhoneNumber),
	}
	indices := svc.indexManager.GetAllIndices()
	for index := range indices {
		if err = indices[index].AddEntity(addToIndexReq); err != nil {
			return err
		}
	}
	return nil
}

func (svc *ContactService) SearchContacts(searchType constants.IndexType, searchString string) ([]*dtos.Contact, error) {
	now := time.Now()
	var res []*dtos.Contact
	searchString = strings.ToLower(searchString)
	contactIds, err := svc.indexManager.GetIndex(searchType).Search(&dtos.SearchIndexRequest{Name: searchString, PhoneNumber: searchString})
	if err != nil {
		return nil, err
	}
	contacts, err := svc.contactRepo.GetContacts(contactIds)
	if err != nil {
		return nil, err
	}
	for _, contact := range contacts {
		res = append(res, mappers.MapContactModelToDTO(contact))
	}
	fmt.Println("Search took ", time.Since(now))
	return res, nil
}
