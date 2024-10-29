package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"sam.com/go/grpc/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {

	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)",
		student.Id, student.Name, student.Age,
	)

	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student

	err := repo.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id).Scan(&student.Id, &student.Name, &student.Age)

	if err != nil {
		return nil, err
	}
	return &student, nil
}
