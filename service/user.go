package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	domain "user/domain"
	repository "user/repository"

	"github.com/google/uuid"
)

type User interface {
	GetUsers() ([]domain.User, error)
	CreateUser(request *http.Request) (*domain.User, error)
}

type UserService struct {
	Repository repository.User
}

func NewUserService(us UserService) User {
	return &UserService{
		Repository: us.Repository,
	}
}

func (service *UserService) GetUsers() ([]domain.User, error) {
	users, err := service.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserService) CreateUser(request *http.Request) (*domain.User, error) {
	var user *domain.User 

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal(requestBody, &user)
	if err != nil  {
		return nil, err
	}
	
	user.ID = uuid.New()

	user, err = service.Repository.Create(user)
	if err != nil  {
		return nil, err
	}

	return user, nil
}