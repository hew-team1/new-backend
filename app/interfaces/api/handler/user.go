package handler

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"

	"guild-hack-api/app/interfaces/api/httputil"
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

func (h *UserHandler) Show(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	strId, ok := vars["user_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	uid, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := h.userUseCase.Show(uid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, user.SwaggerModel(), nil
}
