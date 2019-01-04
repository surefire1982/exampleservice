package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/surefire1982/exampleservice/pkg/entity"
	"github.com/surefire1982/exampleservice/pkg/utils"
)

type TestSuite struct {
	suite.Suite
	svc *Service
}

func (suite *TestSuite) SetupSuite() {
	repo := NewInMemRepository()
	suite.svc = NewService(repo)
}

func (suite *TestSuite) TestStore() {
	t := suite.T()
	user := utils.CreateTestUser()

	id, err := suite.svc.Store(user)

	assert.Nil(t, err)
	assert.NotNil(t, id)

}

func (suite *TestSuite) TestFindAndFindAll() {
	t := suite.T()
	user1 := utils.CreateTestUser()

	user2 := utils.CreateTestUser()

	id1, err1 := suite.svc.Store(user1)
	assert.Nil(t, err1)
	_, err2 := suite.svc.Store(user2)
	assert.Nil(t, err2)

	t.Run("find", func(t *testing.T) {
		u1, err := suite.svc.Find(id1)
		assert.Nil(t, err)
		assert.Equal(t, id1, u1.UserID)
		assert.Equal(t, user1.Username, u1.Username)
	})

	t.Run("find non_existent", func(t *testing.T) {
		u2, err := suite.svc.Find("someID")
		assert.Nil(t, u2)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("find all", func(t *testing.T) {
		users, err := suite.svc.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(users))
	})

}

func (suite *TestSuite) TestDelete() {
	t := suite.T()

	user := utils.CreateTestUser()

	id, err1 := suite.svc.Store(user)
	assert.Nil(t, err1)

	err := suite.svc.Delete(id)
	assert.Nil(t, err)

	findUser, err := suite.svc.Find(id)
	assert.NotNil(t, err)
	assert.Nil(t, findUser)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
