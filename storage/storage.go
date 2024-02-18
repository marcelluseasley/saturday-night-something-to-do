package storage

import (
	"github.com/marcelluseasley/saturday-night-something-to-do/models"
)

type Storager interface {
	CreateUser(user models.User) error
	GetUser(id int) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
	ListUsers() ([]models.User, error)
}
