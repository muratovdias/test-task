package service

import (
	"app2/internal/app/models"
	"app2/internal/app/repository"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

var (
	ErrUserExists = errors.New("User already exists")
)

type User interface {
	CreateUser(models.User) (models.User, error)
	FindUser(string) error
	GetUser(string) (models.User, error)
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) FindUser(email string) error {
	if err := validateEmail(email); err != nil {
		return err
	}
	count, err := s.repo.FindUser(email)
	if count >= 1 && err == nil {
		return ErrUserExists
	} else {
		return err
	}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	if err := validateEmail(user.Email); err != nil {
		return models.User{}, err
	}
	user.Password = generateHash(user.Password, user.Salt)
	user, err := s.repo.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUser(email string) (models.User, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func generateHash(password, salt string) string {
	var str strings.Builder
	str.Grow(len(password) + len(salt))
	str.WriteString(password)
	str.WriteString(salt)
	hash := md5.Sum([]byte(str.String()))
	return hex.EncodeToString(hash[:])
}
