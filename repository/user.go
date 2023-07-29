package repository

import (
	domain "user/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User interface {
	GetAll() ([]domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

type UserDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &UserDB{
		db: db,
	}
}

func(r UserDB) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func(r UserDB) Create(user *domain.User) (*domain.User, error) {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	// Hash the password
	user.Password = hashedPass

	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}