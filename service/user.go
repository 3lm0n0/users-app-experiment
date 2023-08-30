package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	domain "user/domain"
	repository "user/repository"

	"github.com/google/uuid"
)

type User interface {
	GetUsers(ctx context.Context, ids []string) ([]domain.User, error)
	CreateUser(ctx context.Context, request *http.Request) (*domain.User, error)
}

type UserService struct {
	Repository repository.User
}

func NewUserService(us UserService) User {
	return &UserService{
		Repository: us.Repository,
	}
}

func (service *UserService) GetUsers(ctx context.Context, ids []string) ([]domain.User, error) {
	// no ids requested.
	if len(ids) == int(1) && len(ids[0]) == int(0) {
		users, err := service.Repository.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		return users, nil
	}
	// at least 1 ids requested.
	users, err := service.Repository.Get(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserService) CreateUser(ctx context.Context, request *http.Request) (*domain.User, error) {
	var user *domain.User

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()

	user, err = service.Repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
