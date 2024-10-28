package repository

import (
	"context"

	"sam.com/go/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var respImplementation Repository

func SetRepository(repository Repository){
	respImplementation = repository
}

func SetStudent(ctx context.Context student *models.Student) error{
	return respImplementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error){
	return respImplementation.GetStudent(ctx, id)
}