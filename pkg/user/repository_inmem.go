package user

import (
	"time"

	"github.com/rs/xid"
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// IRepo in-memory repository, used for unit testing
type IRepo struct {
	userMap map[string]*entity.User
}

// NewInMemRepository create new in-memory repository implementation
func NewInMemRepository() *IRepo {
	var userMap = map[string]*entity.User{}
	return &IRepo{
		userMap: userMap,
	}
}

// Store a User
func (repo *IRepo) Store(u *entity.User) (string, error) {
	u.UserID = xid.New().String()
	u.CreatedAt = time.Now()
	repo.userMap[u.UserID] = u
	return u.UserID, nil
}

// Find a user
func (repo *IRepo) Find(id string) (*entity.User, error) {
	if repo.userMap[id] == nil {
		return nil, entity.ErrNotFound
	}
	return repo.userMap[id], nil
}

// Delete a user
func (repo *IRepo) Delete(id string) error {
	if repo.userMap[id] == nil {
		return entity.ErrNotFound
	}

	delete(repo.userMap, id)
	return nil
}

// FindAll Users
func (repo *IRepo) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	for _, j := range repo.userMap {
		users = append(users, j)
	}

	return users, nil
}
