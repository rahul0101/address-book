package dtos

import "fmt"

type Contact struct {
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
}

func (contact *Contact) ToString() string {
	return fmt.Sprintf("First Name: %s, LastName: %s, Address: %s, PhoneNumber: %s",
		contact.FirstName, contact.LastName, contact.Address, contact.PhoneNumber)
}
