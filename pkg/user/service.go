package user

import (
	"time"

	"github.com/rs/xid"
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// Service service interface
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Store a user
func (svc *Service) Store(user *entity.User) (string, error) {
	user.UserID = xid.New().String()
	user.CreatedAt = time.Now()
	return svc.repo.Store(user)
}

// Find a user
func (svc *Service) Find(id string) (*entity.User, error) {
	return svc.repo.Find(id)
}

// Delete a user
func (svc *Service) Delete(id string) error {
	return svc.repo.Delete(id)
}

// FindAll users
func (svc *Service) FindAll() ([]*entity.User, error) {
	return svc.repo.FindAll()
}
