package mappers

import (
	"github.com/rahularcota/address-book/dal/model"
	"github.com/rahularcota/address-book/dtos"
)

func MapContactModelToDTO(model *model.Contact) *dtos.Contact {
	return &dtos.Contact{
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		Address:     model.Address,
		PhoneNumber: model.PhoneNumber,
	}
}

func MapContactDTOToModel(dto *dtos.Contact) *model.Contact {
	return &model.Contact{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Address:     dto.Address,
		PhoneNumber: dto.PhoneNumber,
	}
}
