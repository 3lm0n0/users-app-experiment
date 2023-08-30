package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	service "user/service"
)
const usersPath = "/users"

type Handler struct {
	serviceUser service.User
}

func NewUserHandler(su service.User) *Handler {
	return &Handler{
		serviceUser: su,
	}
}

func(h *Handler) Handlers() {
	http.HandleFunc(usersPath, h.handleUsers)
}

func(h *Handler) handleUsers(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("http request method: ", request.Method)
	switch request.Method {
	case http.MethodGet:
		h.handleGetUsers(ctx, response, request)

	case http.MethodPost:
		h.handleCreateUser(ctx, response, request)

	default:
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
}

func(h *Handler) handleGetUsers(ctx context.Context, response http.ResponseWriter, request *http.Request) {
	ids := strings.Split(request.URL.Query().Get("id"), ",")
	dbUsers, err := h.serviceUser.GetUsers(ctx, ids)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusOK, dbUsers)
}

func(h *Handler) handleCreateUser(ctx context.Context, response http.ResponseWriter, request *http.Request) {
	dbUser, err := h.serviceUser.CreateUser(ctx, request)
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