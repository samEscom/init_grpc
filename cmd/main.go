package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sam.com/go/grpc/database"
	"sam.com/go/grpc/server"
	"sam.com/go/grpc/studentpb"
)

func main() {
	list, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	DbString := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"

	repo, err := database.NewPostgresRepository(DbString)

	server := server.NewStdentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
