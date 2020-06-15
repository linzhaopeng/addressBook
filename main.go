package main

import (
	"fmt"
	"linzhaopeng/testGrpc/addressBook"
)

func main() {
	book, fileName := &addressBook.AddressBook{}, "testAddressBook.txt"
	persons := []*addressBook.Person{
		{
			Id:    1,
			Name:  "lzp",
			Email: "linzhaopn@gmail.com",
			PhoneNumbers: []*addressBook.Person_PhoneNumber{
				{
					Number:    "173****4690",
					PhoneType: addressBook.Person_MOBILE,
				},
				{
					Number:    "88****0",
					PhoneType: addressBook.Person_HOME,
				},
			},
		},
		{
			Id:    2,
			Name:  "sunznx",
			Email: "sunznx@gmail.com",
			PhoneNumbers: []*addressBook.Person_PhoneNumber{
				{
					Number:    "157****9661",
					PhoneType: addressBook.Person_MOBILE,
				},
			},
		},
	}

	for _, person := range persons {
		if err := book.AddPerson(person, fileName); err != nil {
			fmt.Println("add person fail:", err)
			return
		}
	}

	book, err := addressBook.GetAddressBook(fileName)
	if err != nil {
		fmt.Println("get addressBook fail:", err)
	}
	book.ListPerson()
}
