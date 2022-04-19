package main

import (
	fmt "fmt"

	pb "github.com/faizainur/learn-grpc/go-addressbook/addressbook_pb"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jogn@doe.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123123", Type: pb.Person_HOME},
		},
	}

	fmt.Println(p.String())
}
