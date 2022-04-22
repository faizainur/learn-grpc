package main

import (
	fmt "fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	"github.com/faizainur/learn-grpc/go-addressbook/addressbook"
)

func main() {
	tcpServer, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Errorf("Error : %q", err.Error())
	}
	s := grpc.NewServer()
	addressbook.RegisterGrpcServer(s, &addressbook.Server{})
	log.Printf("server listening at %v", tcpServer.Addr())
	if err := s.Serve(tcpServer); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
