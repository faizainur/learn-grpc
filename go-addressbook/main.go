package main

import (
	fmt "fmt"
	"log"

	"io/ioutil"

	pb "github.com/faizainur/learn-grpc/go-addressbook/addressbook_pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jogn@doe.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123123", Type: pb.Person_HOME},
		},
	}

	serializedPerson, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("test_pb.txt", serializedPerson, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(person.String())
	fmt.Println("Success")
}
