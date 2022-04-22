package addressbook

import (
	"context"

	pbv2 "github.com/faizainur/learn-grpc/go-addressbook/addressbook_pbv2"
	"google.golang.org/grpc"
)

const (
	Person_HOME   pbv2.Person_PhoneType = pbv2.Person_HOME
	Person_WORK   pbv2.Person_PhoneType = pbv2.Person_WORK
	Person_MOBILE pbv2.Person_PhoneType = pbv2.Person_MOBILE
)

type PhoneNumber = pbv2.Person_PhoneNumber
type Person = pbv2.Person
type AddressBook = pbv2.AddressBook
type Void = pbv2.Void
type StringResponse = pbv2.StringResponse

type Server struct {
	pbv2.UnimplementedAddressBookServiceServer
}

func (s *Server) SayHelloWorld(context.Context, *Void) (*StringResponse, error) {
	return &StringResponse{Message: "HelloWorld"}, nil
}

func (s *Server) GetPerson(context.Context, *Void) (*Person, error) {
	return &Person{
		Name:  "John Doe",
		Id:    1,
		Email: "john@doe.com",
		Phones: []*PhoneNumber{{
			Number: "012391239",
			Type:   Person_HOME,
		}},
	}, nil
}

func NewClient(conn grpc.ClientConnInterface) pbv2.AddressBookServiceClient {
	return pbv2.NewAddressBookServiceClient(conn)
}

func RegisterGrpcServer(grpcServer *grpc.Server, s *Server) {
	pbv2.RegisterAddressBookServiceServer(grpcServer, s)
}

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
