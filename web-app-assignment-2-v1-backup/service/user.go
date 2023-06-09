package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, _ := s.userRepo.GetUserByEmail(user.Email)
	// if err != nil {
	// 	return *user, err
	// }

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	// userData := model.User{}
	userData, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if userData.Email == "" || userData.ID == 0 {
		return nil, errors.New("email not found")
	}

	if userData.Password != user.Password {
		return nil, errors.New("password is incorrect")
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	// Buat claims berisi data username dan role yang akan kita embed ke JWT
	claims := &model.Claims{
		UserID: userData.ID,
		StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// buat token menggunakan library jwt
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenJWT.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
	// TODO: replace this
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	taskCategory, err := s.userRepo.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}
	return taskCategory, nil // TODO: replace this
}
