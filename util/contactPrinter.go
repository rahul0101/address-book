package util

import (
	"fmt"
	"github.com/rahularcota/address-book/dtos"
)

func PrintContacts(contacts []*dtos.Contact) {
	if len(contacts) == 0 {
		fmt.Println("No contacts found!")
		return
	}
	for _, contact := range contacts {
		fmt.Println(contact.ToString())
	}
	fmt.Println()
}
