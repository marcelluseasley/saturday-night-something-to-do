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
	conn, err := ps.dbPool.Acquire(context.Background())
	if err != nil {
		return models.User{}, err
	}
	defer conn.Release()
	var user models.User
	err = conn.QueryRow(context.Background(), "SELECT id, first_name, last_name, email, street_address, city, state, zip_code FROM users WHERE id = $1", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.StreetAddress, &user.City, &user.State, &user.ZipCode)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ps PostgresStorage) UpdateUser(user models.User) error {
	conn, err := ps.dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	cmdTag, err := conn.Exec(context.Background(), "UPDATE users SET first_name = $1, last_name = $2, email = $3, street_address = $4, city = $5, state = $6, zip_code = $7 WHERE id = $8", user.FirstName, user.LastName, user.Email, user.StreetAddress, user.City, user.State, user.ZipCode, user.ID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d", cmdTag.RowsAffected())
	}


	return nil
}

func (ps PostgresStorage) DeleteUser(id int) error {
	conn, err := ps.dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	cmdTag, err := conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d", cmdTag.RowsAffected())
	}
	
	return nil
}

func (ps PostgresStorage) ListUsers() ([]models.User, error) {
	conn, err := ps.dbPool.Acquire(context.Background())
	if err != nil {
		return []models.User{}, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT id, first_name, last_name, email, street_address, city, state, zip_code FROM users")
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.StreetAddress, &user.City, &user.State, &user.ZipCode)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return []models.User{}, err
	}

	return users, nil
}
