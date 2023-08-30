package repository

import (
	"context"
	domain "user/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Get(ctx context.Context, ids []string) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UserDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &UserDB{
		db: db,
	}
}

func(r UserDB) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func(r UserDB) Get(ctx context.Context, ids []string) ([]domain.User, error) {
	var users []domain.User
	var user domain.User
	
	for _, id := range ids {
		uuidParsed, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		user.ID = uuidParsed
		result := r.db.Find(&user)
		if result.Error != nil {
			return nil, result.Error
		}
		users = append(users, user)
	}

	return users, nil
}

func(r UserDB) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
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