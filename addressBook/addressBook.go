package addressBook

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func GetAddressBook(fileName string) (*AddressBook, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	addressBook := &AddressBook{}
	if err = proto.Unmarshal(file, addressBook); err != nil {
		return nil, err
	}
	return addressBook, nil
}

func (b *AddressBook) ListPerson() {
	if b != nil {
		for k, person := range b.Persons {
			fmt.Printf("NO.%d:\n\tPersonID:%d\n", k + 1, person.Id)
			fmt.Printf("\tname:%s\n", person.Name)
			fmt.Printf("\temail:%s\n", person.Email)

			for _, phoneNumber := range person.PhoneNumbers {
				switch phoneNumber.PhoneType {
				case Person_MOBILE:
					fmt.Printf("\tmobile:# %s\n", phoneNumber.Number)
				case Person_HOME:
					fmt.Printf("\thome:# %s\n", phoneNumber.Number)
				case Person_WORK:
					fmt.Printf("\twork:# %s\n", phoneNumber.Number)
				default:
					fmt.Printf("\thas no phone\n")
				}
			}
		}
	}
}

func (b *AddressBook) AddPerson(p *Person, fileName string) error {
	b.Persons = append(b.Persons, p)
	out, err := proto.Marshal(b)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, out, 0644)
}