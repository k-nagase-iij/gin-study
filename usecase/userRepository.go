package usecase

import "github.com/k-nagase-iij/gin-study/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
