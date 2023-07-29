package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	service "user/service"
)

type Handler struct {
	serviceUser service.User
}

func NewUserHandler(su service.User) *Handler {
	return &Handler{
		serviceUser: su,
	}
}

func(h *Handler) Handlers() {
	http.HandleFunc("/users", h.handleUsers)
}

func(h *Handler) handleUsers(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		fmt.Println("http request method: ", request.Method)
		h.handleGetUsers(response, request)

	case http.MethodPost:
		fmt.Println("http request method: ", request.Method)
		h.handleCreateUser(response, request)

	default:
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
}

func(h *Handler) handleGetUsers(response http.ResponseWriter, request *http.Request) {
	dbUsers, err := h.serviceUser.GetUsers()
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusOK, dbUsers)
}

func(h *Handler) handleCreateUser(response http.ResponseWriter, request *http.Request) {
	dbUser, err := h.serviceUser.CreateUser(request)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return
	}

	writeJSONResponse(response, http.StatusCreated, dbUser)
}

func writeJSONResponse(response http.ResponseWriter, status int, value any) error {
	response.WriteHeader(status)
	response.Header().Add("content-type", "application/json; charset=UTF-8")

	return json.NewEncoder(response).Encode(value)
}