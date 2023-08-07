package impl

import (
	"github.com/rahularcota/address-book/dtos"
	"github.com/rahularcota/address-book/indexDataStructures"
	"strings"
)

type NameIndex struct {
	indexDS indexDataStructures.IIndexDS
}

func NewNameIndex(indexDS indexDataStructures.IIndexDS) *NameIndex {
	return &NameIndex{indexDS: indexDS}
}

func (index *NameIndex) Search(req *dtos.SearchIndexRequest) ([]int, error) {
	return index.indexDS.Search(req.Name), nil
}

func (index *NameIndex) AddEntity(req *dtos.AddToIndexRequest) error {
	index.indexDS.Insert(strings.Join([]string{req.FirstName, req.LastName}, " "), req.Id)
	return nil
}
