package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcelluseasley/saturday-night-something-to-do/models"
)

type PostgresStorage struct {
	dbPool *pgxpool.Pool
}

func NewPostgresStorage(dbPool *pgxpool.Pool) PostgresStorage {
	return PostgresStorage{dbPool: dbPool}
}

func (ps PostgresStorage) CreateUser(user models.User) error {

	conn, err := ps.dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	cmdTag, err := conn.Exec(context.Background(), "INSERT INTO users (first_name, last_name, email, street_address, city, state, zip_code) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.FirstName, user.LastName, user.Email, user.StreetAddress, user.City, user.State, user.ZipCode)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d", cmdTag.RowsAffected())
	}
	return nil
}

func (ps PostgresStorage) GetUser(id int) (models.User, error) {

	return models.User{}, nil
}

func (ps PostgresStorage) UpdateUser(user models.User) error {

	return nil
}

func (ps PostgresStorage) DeleteUser(id int) error {

	return nil
}

func (ps PostgresStorage) ListUsers() ([]models.User, error) {

	return []models.User{}, nil
}
