package service

import "user/internal/app/repository"

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
