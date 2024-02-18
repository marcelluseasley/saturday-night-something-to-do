package service

import (
	"github.com/marcelluseasley/saturday-night-something-to-do/models"
	"github.com/marcelluseasley/saturday-night-something-to-do/storage"
)

type UserService struct {
	//db *pgx.Conn
	s storage.Storager
}

func NewUserService(s storage.Storager) UserService {

	return UserService{s: s}
}

func (u *UserService) CreateUser(user models.User) error {

	err := u.s.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
