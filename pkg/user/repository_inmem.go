package user

import (
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// IRepo in memory repository
type IRepo struct {
	userMap map[string]*entity.User
}

// NewInMemRepository create new repository
func NewInMemRepository() *IRepo {
	var userMap = map[string]*entity.User{}
	return &IRepo{
		userMap: userMap,
	}
}

// Store a User
func (repo *IRepo) Store(user *entity.User) (string, error) {
	repo.userMap[user.UserID] = user
	return user.UserID, nil
}
