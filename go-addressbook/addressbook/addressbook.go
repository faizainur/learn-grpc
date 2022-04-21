package addressbook

import (
	pbv2 "github.com/faizainur/learn-grpc/go-addressbook/addressbook_pbv2"
)

const (
	Person_HOME   pbv2.Person_PhoneType = pbv2.Person_HOME
	Person_WORK   pbv2.Person_PhoneType = pbv2.Person_WORK
	Person_MOBILE pbv2.Person_PhoneType = pbv2.Person_MOBILE
)

type PhoneNumber = pbv2.Person_PhoneNumber
type Person = pbv2.Person
type AddressBook = pbv2.AddressBook

func CreatePerson(id int32, name, email string, phones ...*PhoneNumber) *Person {
	return &Person{
		Id:     id,
		Name:   name,
		Email:  email,
		Phones: phones,
	}
}

func CreateAddressBook(persons ...*Person) *AddressBook {
	return &AddressBook{
		Person: persons,
	}
}
