package utils

import (
	"fmt"

	"github.com/rs/xid"
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// CreateTestUser create test user
func CreateTestUser() *entity.User {
	username := xid.New().String()
	email := fmt.Sprintf("%s@email.com", username)
	user := &entity.User{
		Username: username,
		Email:    email,
	}
	return user
}
