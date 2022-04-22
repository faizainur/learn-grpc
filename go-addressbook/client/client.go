package main

import (
	"context"
	"log"
	"time"

	"github.com/faizainur/learn-grpc/go-addressbook/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if err != nil {
			log.Fatal("Error: %v", err)
		}
	}
	defer conn.Close()
	c := addressbook.NewClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHelloWorld(ctx, &addressbook.Void{})
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	p, err := c.GetPerson(ctx, &addressbook.Void{})
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	log.Println(p.String())
}
