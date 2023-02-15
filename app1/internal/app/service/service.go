package service

type Salt interface {
	GetSalt() string
}

type Service struct {
	Salt
}

func NewService() *Service {
	return &Service{
		Salt: NewSaltService(),
	}
}
