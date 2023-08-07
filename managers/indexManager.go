package managers

import (
	"github.com/rahularcota/address-book/constants"
	"github.com/rahularcota/address-book/indices"
)

type IIndexManager interface {
	GetIndex(indexType constants.IndexType) indices.IIndex
	GetAllIndices() []indices.IIndex
}

type IndexManager struct {
	phoneNumberIndex indices.IIndex
	nameIndex        indices.IIndex
}

func NewIndexManager(phoneNumberIndex indices.IIndex, nameIndex indices.IIndex) *IndexManager {
	return &IndexManager{phoneNumberIndex: phoneNumberIndex, nameIndex: nameIndex}
}

func (manager *IndexManager) GetIndex(indexType constants.IndexType) indices.IIndex {
	switch indexType {
	case constants.IndexTypePhoneNumber:
		return manager.phoneNumberIndex
	case constants.IndexTypeName:
		return manager.nameIndex
	}
	return nil
}

func (manager *IndexManager) GetAllIndices() []indices.IIndex {
	return []indices.IIndex{manager.phoneNumberIndex, manager.nameIndex}
}
