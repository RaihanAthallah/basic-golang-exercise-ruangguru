package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	query := fmt.Sprintf("SELECT FROM users WHERE username='%s' AND password='%s'", user.Username, user.Password)
	row := u.db.QueryRow(query)

	var hasil model.User
	err := row.Scan(&hasil.ID, &hasil.Username, &hasil.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
