package user

import (
	"time"

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
	usrRec := user{ID: u.UserID, Username: u.Username, Email: u.Email, CreatedAt: u.CreatedAt}

	err := repo.db.Create(&usrRec).Error
	if err != nil {
		return "", entity.ErrSavingRecord
	}

	return u.UserID, nil
}

// Find a user
func (repo *SQLRepo) Find(id string) (*entity.User, error) {
	var usrRec *user
	repo.db.First(usrRec, id)
	if &usrRec == nil {
		return nil, entity.ErrNotFound
	}

	usrEntity := entity.User{
		UserID:    usrRec.ID,
		Username:  usrRec.Username,
		Email:     usrRec.Email,
		CreatedAt: usrRec.CreatedAt,
	}

	return &usrEntity, nil
}

// Delete a user
func (repo *SQLRepo) Delete(id string) error {
	return nil
}

// FindAll Users
func (repo *SQLRepo) FindAll() ([]*entity.User, error) {
	return nil, nil
}
