package impl

import (
	"github.com/rahularcota/address-book/dtos"
	"github.com/rahularcota/address-book/indexDataStructures"
)

type PhoneNumberIndex struct {
	indexDS indexDataStructures.IIndexDS
}

func NewPhoneNumberIndex(indexDS indexDataStructures.IIndexDS) *PhoneNumberIndex {
	return &PhoneNumberIndex{indexDS: indexDS}
}

func (index *PhoneNumberIndex) Search(req *dtos.SearchIndexRequest) ([]int, error) {
	return index.indexDS.Search(req.PhoneNumber), nil
}

func (index *PhoneNumberIndex) AddEntity(req *dtos.AddToIndexRequest) error {
	index.indexDS.Insert(req.PhoneNumber, req.Id)
	return nil
}
