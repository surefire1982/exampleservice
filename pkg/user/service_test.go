package user

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/surefire1982/exampleservice/pkg/entity"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestSuite struct {
	suite.Suite
	svc *Service
}

func (suite *TestSuite) SetupSuite() {
	t := suite.T()
	db, err := gorm.Open("mysql", "devuser:password@/autobot?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		assert.Fail(t, "Error connecting to database")
	}

	db.DropTableIfExists(&user{})
	repo := NewDBRepository(db)
	suite.svc = NewService(repo)
}

func (suite *TestSuite) TestStore() {
	t := suite.T()
	user := &entity.User{
		Username: "RandomUser",
		Email:    "random_email@email.com",
	}

	id, err := suite.svc.Store(user)

	assert.Nil(t, err)
	assert.NotNil(t, id)

}

func (suite *TestSuite) TestFindAndFindAll() {
	t := suite.T()
	repo := NewInMemRepository()
	svc := NewService(repo)

	user1 := &entity.User{
		Username: "RandomUser1",
		Email:    "random_email1@email.com",
	}

	user2 := &entity.User{
		Username: "RandomUser2",
		Email:    "random_email2@email.com",
	}

	id1, err1 := svc.Store(user1)
	assert.Nil(t, err1)
	_, err2 := svc.Store(user2)
	assert.Nil(t, err2)

	t.Run("find", func(t *testing.T) {
		u1, err := svc.Find(id1)
		assert.Nil(t, err)
		assert.Equal(t, id1, u1.UserID)
		assert.Equal(t, user1.Username, u1.Username)
	})

	t.Run("find non_existent", func(t *testing.T) {
		u2, err := svc.Find("someID")
		assert.Nil(t, u2)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("find all", func(t *testing.T) {
		users, err := svc.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(users))
	})

}

func (suite *TestSuite) TestDelete() {
	t := suite.T()
	repo := NewInMemRepository()
	svc := NewService(repo)

	user1 := &entity.User{
		Username: "RandomUser1",
		Email:    "random_email1@email.com",
	}

	user2 := &entity.User{
		Username: "RandomUser2",
		Email:    "random_email2@email.com",
	}

	id1, err1 := svc.Store(user1)
	assert.Nil(t, err1)
	_, err2 := svc.Store(user2)
	assert.Nil(t, err2)

	err := svc.Delete(id1)
	assert.Nil(t, err)

	users, err := svc.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
