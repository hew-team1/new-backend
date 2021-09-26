package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"guild-hack-api/app/usecase"
	"guild-hack-api/swagger/generated_swagger"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var req generated_swagger.CreateUserRequest
	if err := parseModelRequest(r, &req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := h.userUseCase.Create(
		swag.StringValue(req.UID),
		swag.StringValue(req.Name),
		req.Email.String(),
	)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, user.CreateResponseSwaggerModel(), nil
}

func (h *UserHandler) Index(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	users, err := h.userUseCase.Index()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var resUsers []*generated_swagger.User
	for _, u := range users {
		resUsers = append(resUsers, u.SwaggerModel())
	}
	return http.StatusOK, generated_swagger.UserList{Users: resUsers}, nil
}
