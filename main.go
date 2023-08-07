package main

import (
	"bufio"
	"fmt"
	"github.com/rahularcota/address-book/dal/repo/impl"
	"github.com/rahularcota/address-book/dtos"
	impl4 "github.com/rahularcota/address-book/indexDataStructures/impl"
	impl2 "github.com/rahularcota/address-book/indices/impl"
	"github.com/rahularcota/address-book/managers"
	impl3 "github.com/rahularcota/address-book/service/impl"
	"github.com/rahularcota/address-book/util"
	"os"
)

func main() {

	nameIndex := impl2.NewNameIndex(impl4.NewTrie()) // impl4.NewIndexMap() if prefix search isn't required
	phoneNumberIndex := impl2.NewPhoneNumberIndex(impl4.NewTrie())
	inMemoryContactRepo := impl.NewInMemoryContactRepo()
	indexManager := managers.NewIndexManager(phoneNumberIndex, nameIndex)

	contactService := impl3.NewContactService(inMemoryContactRepo, indexManager)

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Add Contact 2. Search Contacts 3. Exit")
		fmt.Print("Enter choice: ")
		sc.Scan()
		choice := sc.Text()
		switch choice {
		case "1":
			contact := &dtos.Contact{}
			fmt.Print("Enter FirstName: ")
			sc.Scan()
			contact.FirstName = sc.Text()
			fmt.Print("Enter LastName: ")
			sc.Scan()
			contact.LastName = sc.Text()
			fmt.Print("Enter Address: ")
			sc.Scan()
			contact.Address = sc.Text()
			fmt.Print("Enter PhoneNumber: ")
			sc.Scan()
			contact.PhoneNumber = sc.Text()
			err := contactService.AddContact(contact)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Contact added successfully")
			}
		case "2":
			fmt.Println("Search by 1. Name 2. Phone Number")
			fmt.Print("Enter choice: ")
			sc.Scan()
			choice = sc.Text()
			switch choice {
			case "1":
				fmt.Print("Enter Name: ")
				sc.Scan()
				choice = sc.Text()
				contacts, err := contactService.SearchContacts("NAME", choice)
				if err != nil {
					fmt.Println(err.Error())
				}
				util.PrintContacts(contacts)
			case "2":
				fmt.Print("Enter PhoneNumber: ")
				sc.Scan()
				choice = sc.Text()
				contacts, err := contactService.SearchContacts("PHONE_NUMBER", choice)
				if err != nil {
					fmt.Println(err.Error())
				}
				util.PrintContacts(contacts)
			default:
				fmt.Println("Invalid option")
			}
		case "3":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
