package user

import (
	"time"

	"github.com/rs/xid"

	"github.com/jinzhu/gorm"
	"github.com/surefire1982/exampleservice/pkg/entity"
)

// SQLRepo database implementation of repository
type SQLRepo struct {
	db *gorm.DB
}

// unexported struct used only for the repository
type user struct {
	ID        string // gorm automatically marks as primary key
	Username  string `gorm:"UNIQUE"`
	Email     string `gorm:"UNIQUE"`
	CreatedAt time.Time
}

// NewDBRepository creates new database repository implementation
func NewDBRepository(db *gorm.DB) *SQLRepo {
	// Migrate the schema
	// Add table suffix when create tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user{})

	return &SQLRepo{
		db: db,
	}
}

// Store a User
func (repo *SQLRepo) Store(u *entity.User) (string, error) {
	uid := xid.New().String()
	usrRec := user{ID: uid, Username: u.Username, Email: u.Email, CreatedAt: u.CreatedAt}

	err := repo.db.Create(&usrRec).Error
	if err != nil {
		return "", err
	}

	return uid, nil
}

// Find a user
func (repo *SQLRepo) Find(id string) (*entity.User, error) {
	findUsr := &user{
		ID: id,
	}
	err := repo.db.First(findUsr).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, entity.ErrNotFound
		}
		// otherwise just return error
		return nil, err
	}

	usrEntity := entity.User{
		UserID:    findUsr.ID,
		Username:  findUsr.Username,
		Email:     findUsr.Email,
		CreatedAt: findUsr.CreatedAt,
	}

	return &usrEntity, nil
}

// Delete a user
func (repo *SQLRepo) Delete(id string) error {
	delUsr := &user{
		ID: id,
	}
	err := repo.db.Delete(delUsr).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll Users
func (repo *SQLRepo) FindAll() ([]*entity.User, error) {
	var users []user
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	var entityUsers []*entity.User

	for _, user := range users {
		entityUser := &entity.User{
			UserID:    user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
		entityUsers = append(entityUsers, entityUser)
	}
	return entityUsers, nil
}
