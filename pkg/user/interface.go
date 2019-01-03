package user

import (
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// Repository repository interface for users
type Repository interface {
	// likely change from string to a specific ID as needed
	Store(u *entity.User) (string, error)
	Find(id string) (*entity.User, error)
	Delete(id string) error
	FindAll() ([]*entity.User, error)
}
