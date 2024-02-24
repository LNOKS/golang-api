package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepository interface {
	GetAll(searchTerm string) ([]User, error)
	GetByUsername(username string) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int) error
	DeleteBatch(ids []int) interface{}
}

type userRepository struct {
	db *sqlx.DB
}

func (m userRepository) GetByUsername(username string) (User, error) {
	var user User
	err := m.db.Get(&user, "SELECT * FROM users WHERE user_name = $1", username)
	if err != nil {
		return User{}, errors.Wrap(err, "unable to fetch user by username")
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

func (m userRepository) GetAll(searchTerm string) ([]User, error) {
	var users []User
	err := m.db.Select(&users, "SELECT * FROM users WHERE first_name LIKE $1 OR last_name LIKE $1", "%"+searchTerm+"%")
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch all users")
	}
	return users, nil
}

func (m userRepository) Create(user User) error {
	_, err := m.db.NamedExec(`INSERT INTO users (first_name, last_name, user_name, password) VALUES (:first_name, :last_name, :user_name, :password)`, map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"user_name":  user.Username,
		"password":   user.Password,
	})

	if err != nil {
		return errors.Wrap(err, "unable to create user")
	}
	return nil
}

func (m userRepository) Update(user User) error {
	_, err := m.db.NamedExec(`UPDATE users SET first_name = :first_name, last_name = :last_name, user_name = :user_name, password = :password, birthdate = :birthdate, work_start_date = :work_start_date WHERE id = :id`, map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"user_name":  user.Username,
		"password":   user.Password,
		"id":         user.ID,
	})
	if err != nil {
		return errors.Wrap(err, "unable to update user")
	}
	return nil
}

func (m userRepository) Delete(id int) error {
	_, err := m.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return errors.Wrap(err, "unable to delete user")
	}
	return nil
}

func (m userRepository) DeleteBatch(ids []int) interface{} {
	query, args, err := sqlx.In("DELETE FROM users WHERE id IN (?)", ids)
	if err != nil {
		return errors.Wrap(err, "unable to delete users")
	}
	query = m.db.Rebind(query)
	_, err = m.db.Exec(query, args...)
	if err != nil {
		return errors.Wrap(err, "unable to delete users")
	}
	return nil
}
