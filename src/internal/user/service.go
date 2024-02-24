package user

import (
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"hrm-api/src/internal/utils"
)

type UserService interface {
	GetAll(searchTerm string) ([]User, error)
	GetByUsername(username string) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int) error
	DeleteBatch(ids []int) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (m userService) GetAll(searchTerm string) ([]User, error) {
	users, err := m.repo.GetAll(searchTerm)
	if err != nil {
		log.Error("Error getting all users: ", err)
	}
	if users == nil {
		return []User{}, nil
	}
	return users, nil
}

func (m userService) GetByUsername(username string) (User, error) {
	user, err := m.repo.GetByUsername(username)
	if err != nil {
		log.Error("Error getting user by username: ", err)
		return User{}, errors.Wrap(err, "User not found")
	}
	return user, nil
}

func (m userService) Create(user User) error {
	user.Password = utils.HashPassword(user.Password)
	err := m.repo.Create(user)
	if err != nil {
		log.Error("Error creating user: ", err)
	}
	return nil
}

func (m userService) Update(user User) error {
	if user.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}
	err := m.repo.Update(user)
	if err != nil {
		log.Error("Error updating user: ", err)
	}
	return nil
}

func (m userService) Delete(id int) error {
	err := m.repo.Delete(id)
	if err != nil {
		log.Error("Error deleting user: ", err)
	}
	return nil
}

func (m userService) DeleteBatch(ids []int) error {
	err := m.repo.DeleteBatch(ids)
	if err != nil {
		log.Error("Error deleting users: ", err)
	}
	return nil
}
