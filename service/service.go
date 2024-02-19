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

func (u *UserService) GetUser(id int) (models.User, error) {
	user, err := u.s.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(user models.User) error {
	err := u.s.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(id int) error {
	err := u.s.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ListUsers() ([]models.User, error) {
	users, err := u.s.ListUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}


