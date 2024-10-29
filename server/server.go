package server

import (
	"context"

	"sam.com/go/grpc/models"
	"sam.com/go/grpc/repository"
	"sam.com/go/grpc/studentpb"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStdentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (server *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := server.repo.GetStudent(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (server *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudetResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := server.repo.SetStudent(ctx, student)

	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudetResponse{
		Id: student.Id,
	}, nil
}
